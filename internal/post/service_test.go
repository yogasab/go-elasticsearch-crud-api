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

	postStorage, err := elasticsearch.NewPostStorage(*elastic)
	postService := NewPostService(postStorage)

	var req InsertDocumentRequest
	req.Title = "Post from service"
	req.Text = "Text from service"
	req.Tags = []string{"tags8", "tags9"}
	newPost, errRest := postService.InsertDocument(ctx, req)

	assert.Nil(t, err)
	assert.Nil(t, errRest)
	assert.NotNil(t, postStorage)
	assert.NotNil(t, postService)
	assert.NotNil(t, newPost)

	assert.Equal(t, req.Title, newPost.Title)
	assert.Equal(t, req.Text, newPost.Text)
	assert.Equal(t, req.Tags, newPost.Tags)
}

func TestErrorFindDocumentByID(t *testing.T) {
	elastic := newPostService()
	postStorage, err := elasticsearch.NewPostStorage(*elastic)
	postService := NewPostService(postStorage)

	unavailableDocumentID := "5d16a5b5-2e4d-4de2-baba-3d12bc69b2a5"
	currentPost, errRest := postService.FindDocumentByID(ctx, unavailableDocumentID)

	assert.Nil(t, err)
	assert.Nil(t, currentPost)
	assert.NotNil(t, errRest)
	assert.NotNil(t, errRest.Code())
	assert.NotNil(t, errRest.Status())
	assert.NotNil(t, errRest.Message())
}

func TestFindDocumentByID(t *testing.T) {
	elastic := newPostService()
	err := elastic.CreateIndex("posts")

	postStorage, err := elasticsearch.NewPostStorage(*elastic)
	postService := NewPostService(postStorage)

	currentPost, errRest := postService.FindDocumentByID(ctx, "1c1802cd-a99e-4bd6-92f8-33213ec29ed9")

	assert.Nil(t, err)
	assert.Nil(t, errRest)
	assert.NotNil(t, currentPost)

	assert.Equal(t, "New Title from Test", currentPost.Title)
	assert.Equal(t, "New Text from Test", currentPost.Text)
	assert.Equal(t, []string{"tags1", "tags2"}, currentPost.Tags)
	assert.Equal(t, "2022-08-10 12:51:08.5610373 +0000 UTC", currentPost.CreatedAt.String())
}

func TestDeleteDocumentByID(t *testing.T) {
	elastic := newPostService()
	err := elastic.CreateIndex("posts")

	postStorage, err := elasticsearch.NewPostStorage(*elastic)
	postService := NewPostService(postStorage)

	documentID := "new id insert here"
	currentPost, errRest := postService.DeleteDocumentByID(ctx, documentID)

	assert.Nil(t, err)
	assert.Nil(t, errRest)
	assert.NotNil(t, currentPost)

	assert.EqualValues(t, true, currentPost)
}

func TestErrorDeleteDocumentByID(t *testing.T) {
	elastic := newPostService()
	err := elastic.CreateIndex("posts")

	postStorage, err := elasticsearch.NewPostStorage(*elastic)
	postService := NewPostService(postStorage)

	documentID := "5f821657-1dea-492c-897b-416695e5f1c9"
	currentPost, errRest := postService.DeleteDocumentByID(ctx, documentID)

	assert.Nil(t, err)
	assert.NotNil(t, errRest)
	assert.NotNil(t, currentPost)
	assert.NotNil(t, errRest.Code())
	assert.NotNil(t, errRest.Status())
	assert.NotNil(t, errRest.Message())
	assert.EqualValues(t, false, currentPost)
}

func TestUpdateDocumentByID(t *testing.T) {
	elastic := newPostService()
	err := elastic.CreateIndex("posts")

	postStorage, err := elasticsearch.NewPostStorage(*elastic)
	postService := NewPostService(postStorage)

	var updatedReq UpdateDocumentRequest
	updatedReq.ID = "1c1802cd-a99e-4bd6-92f8-33213ec29ed9"
	updatedReq.Title = "Post from service updated"
	updatedReq.Text = "Text from service updated"
	updatedReq.Tags = []string{"tags8", "tags9"}
	updatedPost, errRest := postService.UpdateDocumentByID(ctx, updatedReq)

	assert.Nil(t, err)
	assert.Nil(t, errRest)
	assert.NotNil(t, updatedPost)

	assert.EqualValues(t, true, updatedPost)
}

func TestErrorNotFoundUpdateDocumentByID(t *testing.T) {
	elastic := newPostService()
	err := elastic.CreateIndex("posts")

	postStorage, err := elasticsearch.NewPostStorage(*elastic)
	postService := NewPostService(postStorage)

	var updatedReq UpdateDocumentRequest
	updatedReq.ID = "1c1802cd-a99e-4bd6-92f8"
	updatedReq.Title = "Post from service updated"
	updatedReq.Text = "Text from service updated"
	updatedReq.Tags = []string{"tags8", "tags9"}
	updatedPost, errRest := postService.UpdateDocumentByID(ctx, updatedReq)

	assert.Nil(t, err)
	assert.NotNil(t, errRest)
	assert.NotNil(t, updatedPost)

	assert.EqualValues(t, false, updatedPost)
	assert.EqualValues(t, "failed", errRest.Status())
	assert.EqualValues(t, "error document not found", errRest.Message())
}
