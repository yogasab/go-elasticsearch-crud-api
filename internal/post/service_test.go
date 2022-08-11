package post

import (
	"context"
	"log"
	"net/http"
	"testing"
	"yogasab/go-elasticsearch-crud-api/internal/pkg/storage/elasticsearch"

	"github.com/stretchr/testify/assert"
)

var (
	ctx = context.Background()
)

func newPostService() *elasticsearch.Elasticsearch {
	elastic, err := elasticsearch.New([]string{"http://127.0.0.1:9206"})
	if err != nil {
		log.Fatalln(err)
	}
	return elastic
}

func TestErrorInsertDocument(t *testing.T) {
	elastic := newPostService()
	postStorage, err := elasticsearch.NewPostStorage(*elastic)
	postService := NewPostService(postStorage)

	var req InsertDocumentRequest
	req.Title = "Post from service"
	req.Text = "Text from service"
	req.Tags = []string{"tags8", "tags9"}
	newPost, errRest := postService.InsertDocument(ctx, req)

	assert.Nil(t, err)
	assert.NotNil(t, errRest)
	assert.Nil(t, newPost)
	assert.NotNil(t, postStorage)
	assert.NotNil(t, postService)

	assert.Equal(t, http.StatusInternalServerError, errRest.Code())
	assert.Equal(t, "error", errRest.Status())
	assert.Equal(t, "error while inserting new document", errRest.Message())
}

func TestInsertDocument(t *testing.T) {
	elastic := newPostService()
	err := elastic.CreateIndex("posts")

	postStorage, errRest := elasticsearch.NewPostStorage(*elastic)
	postService := NewPostService(postStorage)

	var req InsertDocumentRequest
	req.Title = "Post from service"
	req.Text = "Text from service"
	req.Tags = []string{"tags8", "tags9"}
	newPost, err := postService.InsertDocument(ctx, req)

	assert.Nil(t, err)
	assert.Nil(t, errRest)
	assert.NotNil(t, postStorage)
	assert.NotNil(t, postService)
	assert.NotNil(t, newPost)

	assert.Equal(t, req.Title, newPost.Title)
	assert.Equal(t, req.Text, newPost.Text)
	assert.Equal(t, req.Tags, newPost.Tags)
}
