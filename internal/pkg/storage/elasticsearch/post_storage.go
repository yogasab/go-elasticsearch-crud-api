package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"
	"yogasab/go-elasticsearch-crud-api/internal/pkg/storage"

	"github.com/elastic/go-elasticsearch/v8/esapi"
)

var PostStorage storage.PostStore = postStorage{}

type postStorage struct {
	elastic Elasticsearch
	timeout time.Duration
}

func NewPostStorage(elastic Elasticsearch) (postStorage, error) {
	return postStorage{
		elastic: elastic,
		timeout: time.Second * 10,
	}, nil
}

func (p postStorage) Insert(ctx context.Context, post storage.Post) error {
	body, err := json.Marshal(post)
	if err != nil {
		return errors.New(fmt.Sprintf("error while marshalling body json %v", err))
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
		return errors.New(fmt.Sprintf("error inserting %v", err))
	}
	defer res.Body.Close()
	if res.StatusCode == 409 {
		return errors.New(fmt.Sprintf("conflict %v", err))
	}
	if res.IsError() {
		return errors.New(fmt.Sprintf("error while inserting %v", err))
	}

	return nil
}
