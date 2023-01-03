package items

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/PaulTabaco/bookstore_items-api/clients/elasticsearch"
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

	if result.IsError() {
		return rest_errors.NewInternalServerError("error when trying to save item", errors.New("database error"))
	}

	// Deserialize the response into a map.
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
