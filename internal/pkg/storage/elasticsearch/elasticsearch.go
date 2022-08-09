package elasticsearch

import (
	"errors"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
)

type Elasticsearch struct {
	client *elasticsearch.Client
	index  string
	alias  string
}

type document struct {
	Source interface{} `json:"_source"`
}

func New(address []string) (*Elasticsearch, error) {
	cfg := elasticsearch.Config{
		Addresses: address,
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return &Elasticsearch{
		client: client,
	}, nil
}

func (e *Elasticsearch) CreateIndex(index string) error {
	e.index = index
	e.alias = index + "_alias"

	res, err := e.client.Indices.Exists([]string{e.index})
	if err != nil {
		return errors.New("error while checking the exisiting index")
	}
	if res.StatusCode == 200 {
		return nil
	}
	if res.StatusCode != 404 {
		return errors.New(fmt.Sprintf("index not found %s", res.String()))
	}

	res, err = e.client.Indices.Create(e.index)
	if err != nil {
		return errors.New(fmt.Sprintf("error while create new index %s", err.Error()))
	}
	if res.IsError() {
		return errors.New(fmt.Sprintln("error in index creation %s", res.String()))
	}

	res, err = e.client.Indices.PutAlias([]string{e.index}, e.alias)
	if err != nil {
		return errors.New(fmt.Sprintf("cannot create index %v", err))
	}
	if res.IsError() {
		return errors.New(fmt.Sprintf("error while naming alias %s", res.String()))
	}

	return nil
}
