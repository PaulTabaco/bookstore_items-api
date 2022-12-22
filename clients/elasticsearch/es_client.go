package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	// "github.com/elastic/go-elasticsearch/v8"

	"github.com/PaulTabaco/bookstore_utils/logger"
	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v8"
)

var (
	Client esClientInerface = &esClient{}

	cfg = elasticsearch.Config{
		Addresses: []string{
			"https://localhost:9200",
			// "https://localhost:9201",
		},
		Username: "foo",
		Password: "bar",
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
	setClient(*elasticsearch.Client)
	Index(string, interface{}) (*esapi.Response, error)
}

type esClient struct {
	client *elasticsearch.Client
}

func Init() {
	// log := logger.GetLogger()
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	Client.setClient(es)

	// res, err := es.Info()
	// if err != nil {
	// 	log.Fatalf("Error getting response: %s", err)
	// }

	// defer res.Body.Close()
	// log.Println(res)
}

func (c *esClient) Index(index string, doc interface{}) (*esapi.Response, error) {
	// data, err := json.Marshal(struct{ Title string }{Title: "title"})
	data, err := json.Marshal(doc)
	if err != nil {

		logger.Error("error marshaling document", err)
		return nil, err
	}

	req := esapi.IndexRequest{
		Index: index,
		// DocumentID: strconv.Itoa(i + 1),
		Body:    bytes.NewReader(data),
		Refresh: "true",
	}

	res, err := req.Do(context.Background(), c.client)
	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to index document in index %s", index), err)
		return nil, err
	}
	defer res.Body.Close()

	return res, nil
}

func (c *esClient) setClient(client *elasticsearch.Client) {
	c.client = client
}
