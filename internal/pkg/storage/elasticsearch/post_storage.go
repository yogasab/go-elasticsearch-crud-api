package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"
	"yogasab/go-elasticsearch-crud-api/internal/pkg/storage"
	"yogasab/go-elasticsearch-crud-api/internal/utils/http_errors"

	"github.com/elastic/go-elasticsearch/v8/esapi"
)

var PostStorage storage.PostStore = postStorage{}

type postStorage struct {
	elastic Elasticsearch
	timeout time.Duration
}

func NewPostStorage(elastic Elasticsearch) (postStorage, http_errors.RestErrors) {
	return postStorage{
		elastic: elastic,
		timeout: time.Second * 10,
	}, nil
}

func (p postStorage) Insert(ctx context.Context, post storage.Post) http_errors.RestErrors {
	body, err := json.Marshal(post)
	if err != nil {
		return http_errors.NewBadRequestError("error while marshalling body json", []interface{}{err})
	}
	req := esapi.CreateRequest{
		Index:      p.elastic.alias,
		DocumentID: post.ID,
		Body:       bytes.NewReader(body),
	}
	ctx, cancel := context.WithTimeout(ctx, p.timeout)
	defer cancel()
	res, err := req.Do(ctx, p.elastic.client)
	if err != nil {
		return http_errors.NewInternalServerError("error while inserting new document", []interface{}{err})
	}
	defer res.Body.Close()
	if res.StatusCode == 409 {
		return http_errors.NewInternalServerError("error while conflict", []interface{}{err})
	}
	if res.IsError() {
		return http_errors.NewInternalServerError("error while inserting new document", []interface{}{err})
	}

	return nil
}

func (p postStorage) FindByID(ctx context.Context, ID string) (*storage.Post, http_errors.RestErrors) {
	req := esapi.GetRequest{
		Index:      p.elastic.alias,
		DocumentID: ID,
	}
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	res, err := req.Do(ctx, p.elastic.client)
	if err != nil {
		return nil, http_errors.NewInternalServerError("error while find one document", []interface{}{err})
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusNotFound {
		return nil, http_errors.NewStatusNotFoundError("error document not found", []interface{}{err})
	}
	if res.IsError() {
		return nil, http_errors.NewInternalServerError("error while inserting new document", []interface{}{err})
	}
	var (
		post storage.Post
		body document
	)
	body.Source = &post
	if err = json.NewDecoder(res.Body).Decode(&body); err != nil {
		return nil, http_errors.NewBadRequestError("error while inserting new document", []interface{}{err})
	}
	return &post, nil
}
