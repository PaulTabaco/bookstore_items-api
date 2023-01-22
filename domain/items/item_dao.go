package items

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/PaulTabaco/bookstore_items-api/clients/elasticsearch"
	"github.com/PaulTabaco/bookstore_items-api/domain/queries"
	"github.com/PaulTabaco/bookstore_utils/logger"
	"github.com/PaulTabaco/bookstore_utils/rest_errors"
)

const (
	indexItems = "items"
)

func (i *Item) Save() rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save item", errors.New("database error"))
	}

	var r map[string]interface{}
	if err := json.NewDecoder(result.Body).Decode(&r); err != nil {
		logger.GetLogger().Printf("Error parsing the response body: %s", err)
		return rest_errors.NewInternalServerError("error when trying decode index response to item", errors.New(" error"))
	}

	id := fmt.Sprintf("%s", r["_id"])
	if id == "" {
		return rest_errors.NewInternalServerError("error when trying get id from response item", errors.New(" error"))
	}

	i.Id = id

	return nil
}

func (i *Item) Get() rest_errors.RestErr {
	itemId := i.Id
	result, err := elasticsearch.Client.Get(indexItems, i.Id)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to get item", errors.New("database error"))
	}

	if result.StatusCode == 404 {
		return rest_errors.NewNotFoundError(fmt.Sprintf("no item fount with id: %s", i.Id))
	}

	// Deserialize the response into a map.
	var r map[string]interface{}
	if err := json.NewDecoder(result.Body).Decode(&r); err != nil {
		logger.GetLogger().Printf("Error parsing the response body: %s", err)
		return rest_errors.NewInternalServerError("error when trying decode get response to item", errors.New(" error"))
	}

	source := (r["_source"])
	bytes, _ := json.Marshal(source)

	if err := json.Unmarshal(bytes, &i); err != nil {
		return rest_errors.NewInternalServerError("error when trying to parse database response", errors.New("database error"))
	}
	i.Id = itemId
	return nil
}

func (i *Item) Search(query queries.EsQuery) ([]Item, rest_errors.RestErr) {
	finalQuery := query.ToCorrectEsQuery()
	var bytes bytes.Buffer
	if err := json.NewEncoder(&bytes).Encode(finalQuery); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	result, err := elasticsearch.Client.Search(indexItems, bytes)
	if err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to search documents", errors.New("database error"))
	}
	fmt.Println(result.String())

	if result.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(result.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				result.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	var r map[string]interface{}

	if err := json.NewDecoder(result.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	items := make([]Item, 0)
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {

		id := hit.(map[string]interface{})["_id"]
		fmt.Println(id)
		source := hit.(map[string]interface{})["_source"]
		fmt.Println(source)
		bytes, _ := json.Marshal(source)

		var item Item

		if err := json.Unmarshal(bytes, &item); err != nil {
			return nil, rest_errors.NewInternalServerError("error when trying parse search response", errors.New("database error"))
		}

		items = append(items, item)

		log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	}

	if len(items) == 0 {
		return nil, rest_errors.NewNotFoundError("no items found by given criterias ")
	}

	return items, nil
}
