package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/PaulTabaco/bookstore_utils/logger"

	esv8 "github.com/elastic/go-elasticsearch/v8"
	esv8api "github.com/elastic/go-elasticsearch/v8/esapi"
)

var (
	Client esClientInerface = &esClient{}

	cfg = esv8.Config{
		Addresses: []string{
			"http://localhost:9200",
			// "https://localhost:9201",
		},
		// Username: "foo",
		// Password: "bar",
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			// TLSClientConfig: &tls.Config{
			// 	MinVersion: tls.VersionTLS12,
			// 	// ...
			// },
			// ...
		},
	}
)

type esClientInerface interface {
	setClient(*esv8.Client)
	Index(string, interface{}) (*esv8api.Response, error)
	Get(string, string) (*esv8api.Response, error)
}

type esClient struct {
	//client *elasticsearch.Client
	client *esv8.Client
}

func Init() {
	// log := logger.GetLogger()
	es, err := esv8.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	Client.setClient(es)

	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()
	log.Println(res)
}

func (c *esClient) setClient(client *esv8.Client) {
	c.client = client
}

func (c *esClient) Index(index string, doc interface{}) (*esv8api.Response, error) {
	data, err := json.Marshal(doc)
	if err != nil {
		logger.Error("error marshaling document", err)
		return nil, err
	}

	req := esv8api.IndexRequest{
		Index:   index,
		Body:    bytes.NewReader(data),
		Refresh: "true",
	}

	res, err := req.Do(context.Background(), c.client)
	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to index document in index %s", index), err)
		return nil, err
	}

	return res, nil
}

func (c *esClient) Get(index string, id string) (*esv8api.Response, error) {
	resp, err := c.client.Get(index, id)
	// fmt.Println(resp.String())

	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to get document by id - %s", id), err)
		return nil, err
	}

	return resp, nil

}
