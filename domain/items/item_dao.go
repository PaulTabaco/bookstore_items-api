package items

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/PaulTabaco/bookstore_items-api/clients/elasticsearch"
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
	// TODO:- change logic to evoid Second check for error
	if result.IsError() {
		log.Printf("[%s] Error indexing document ID=%d", result.Status(), 1)
		return rest_errors.NewInternalServerError("error when trying to save item", errors.New("database error"))
	}

	// Deserialize the response into a map.
	var r map[string]interface{}
	if err := json.NewDecoder(result.Body).Decode(&r); err != nil {
		log.Printf("Error parsing the response body: %s", err)
		return rest_errors.NewInternalServerError("error when trying decode index response to item", errors.New(" error"))
	}

	// Print the response status and indexed document version.
	log.Printf("[%s] %s; version=%d", result.Status(), r["result"], int(r["_version"].(float64)))

	// TODO: - Check it
	id := fmt.Sprintf("%s", r["_id"])
	if id != "" {
		return rest_errors.NewInternalServerError("error when trying get id from response item", errors.New(" error"))
	}

	i.Id = id

	return nil
}
