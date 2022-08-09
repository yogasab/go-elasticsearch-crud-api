package elasticsearch

import (
	"yogasab/go-elasticsearch-crud-api/internal/utils/http_errors"

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

func New(address []string) (*Elasticsearch, http_errors.RestErrors) {
	cfg := elasticsearch.Config{
		Addresses: address,
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, http_errors.NewInternalServerError("cannot create new elasticsearch client", []interface{}{err})

	}
	return &Elasticsearch{
		client: client,
	}, nil
}

func (e *Elasticsearch) CreateIndex(index string) http_errors.RestErrors {
	e.index = index
	e.alias = index + "_alias"

	res, err := e.client.Indices.Exists([]string{e.index})
	if err != nil {
		return http_errors.NewInternalServerError("cannot check index existence", []interface{}{err})
	}
	if res.StatusCode == 200 {
		return nil
	}
	if res.StatusCode != 404 {
		return http_errors.NewStatusNotFoundError("error in index existence response:", []interface{}{res.String()})
	}

	res, err = e.client.Indices.Create(e.index)
	if err != nil {
		return http_errors.NewInternalServerError("cannot create index", []interface{}{err})
	}
	if res.IsError() {
		return http_errors.NewInternalServerError("error in index creation response:", []interface{}{res.String()})
	}

	res, err = e.client.Indices.PutAlias([]string{e.index}, e.alias)
	if err != nil {
		return http_errors.NewInternalServerError("cannot create index alias", []interface{}{err})
	}
	if res.IsError() {
		return http_errors.NewInternalServerError("error in index alias creation response:", []interface{}{res.String()})
	}

	return nil
}
