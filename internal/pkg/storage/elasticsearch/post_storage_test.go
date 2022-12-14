package elasticsearch

import (
	"context"
	"net/http"
	"testing"
	"time"
	"yogasab/go-elasticsearch-crud-api/internal/pkg/storage"

	"github.com/stretchr/testify/assert"
)

var (
	ctx        = context.Background()
	created_at = time.Now().UTC()
	newPost    = storage.Post{
		ID:        id,
		Title:     "New Title from Test",
		Text:      "New Text from Test",
		Tags:      []string{"tags1", "tags2"},
		CreatedAt: &created_at,
	}
	updatedPost = storage.Post{
		ID:        id,
		Title:     "New Title from Update",
		Text:      "New Text from Update",
		Tags:      []string{"tags2", "tags3"},
		CreatedAt: &created_at,
	}
)

// ===============================================
// Change to wrong elastic address
func TestErrorInsertNewDocument(t *testing.T) {
	elastic, err := New([]string{wrongAddress})
	elastic.index = "posts"
	elastic.alias = elastic.index + "_alias"

	ps, errRest := NewPostStorage(*elastic)
	errRest = ps.Insert(ctx, newPost)

	assert.Nil(t, err)
	assert.NotNil(t, errRest)
	assert.NotNil(t, ps)

	assert.EqualValues(t, http.StatusInternalServerError, errRest.Code())
	assert.EqualValues(t, "error", errRest.Status())
	assert.EqualValues(t, "error while inserting new document", errRest.Message())
}
func TestResErrorInsertNewDocument(t *testing.T) {
	elastic, err := New([]string{address})

	ps, errRest := NewPostStorage(*elastic)
	errRest = ps.Insert(ctx, newPost)

	assert.Nil(t, err)
	assert.NotNil(t, errRest)
	assert.NotNil(t, ps)

	assert.EqualValues(t, http.StatusInternalServerError, errRest.Code())
	assert.EqualValues(t, "error", errRest.Status())
	assert.EqualValues(t, "error while inserting new document", errRest.Message())
}
func TestInsertNewDocument(t *testing.T) {
	elastic, err := New([]string{address})
	elastic.index = "posts"
	elastic.alias = elastic.index + "_alias"

	ps, errRest := NewPostStorage(*elastic)
	errRest = ps.Insert(ctx, newPost)

	assert.Nil(t, err)
	assert.Nil(t, errRest)
	assert.NotNil(t, ps)
}

// ===============================================

// ===============================================
func TestErrFindByID(t *testing.T) {
	elastic, err := New([]string{wrongAddress})
	elastic.index = "posts"
	elastic.alias = elastic.index + "_alias"

	ps, errRest := NewPostStorage(*elastic)
	currentPost, errRest := ps.FindByID(ctx, id)

	assert.NotNil(t, elastic)
	assert.Nil(t, err)
	assert.NotNil(t, errRest)
	assert.Nil(t, currentPost)

	assert.EqualValues(t, http.StatusInternalServerError, errRest.Code())
	assert.EqualValues(t, "error", errRest.Status())
	assert.EqualValues(t, "error while find one document", errRest.Message())
}
func TestErrNotFoundFindByID(t *testing.T) {
	elastic, err := New([]string{address})
	elastic.index = "posts"
	elastic.alias = elastic.index + "_alias"

	ps, errRest := NewPostStorage(*elastic)
	currentPost, errRest := ps.FindByID(ctx, "710fd955-c2b8-4451-9ade-1cd8055d0dbe")

	assert.Nil(t, err)
	assert.Nil(t, currentPost)
	assert.NotNil(t, elastic)
	assert.NotNil(t, errRest)

	assert.EqualValues(t, http.StatusNotFound, errRest.Code())
	assert.EqualValues(t, "failed", errRest.Status())
	assert.EqualValues(t, "error document not found", errRest.Message())
}
func TestFoundDocumentFindByID(t *testing.T) {
	elastic, err := New([]string{address})
	elastic.index = "posts"
	elastic.alias = elastic.index + "_alias"

	ps, errRest := NewPostStorage(*elastic)
	currentPost, errRest := ps.FindByID(ctx, "1c1802cd-a99e-4bd6-92f8-33213ec29ed9")

	assert.Nil(t, err)
	assert.Nil(t, errRest)
	assert.NotNil(t, currentPost)
	assert.NotNil(t, elastic)

	assert.EqualValues(t, "1c1802cd-a99e-4bd6-92f8-33213ec29ed9", currentPost.ID)
	assert.EqualValues(t, newPost.Title, currentPost.Title)
	assert.EqualValues(t, newPost.Text, currentPost.Text)
	assert.EqualValues(t, newPost.Tags, currentPost.Tags)
	assert.EqualValues(t, "2022-08-10 12:51:08.5610373 +0000 UTC", currentPost.CreatedAt.String())
}

// ===============================================

// ===============================================
func TestErrDeleteByID(t *testing.T) {
	elastic, err := New([]string{wrongAddress})
	elastic.index = "posts"
	elastic.alias = elastic.index + "_alias"

	ps, errRest := NewPostStorage(*elastic)
	errRest = ps.DeleteByID(ctx, "710fd955-c2b8-4451-9ade-1cd8055d0dbe")

	assert.Nil(t, err)
	assert.NotNil(t, elastic)
	assert.NotNil(t, errRest)

	assert.EqualValues(t, http.StatusInternalServerError, errRest.Code())
	assert.EqualValues(t, "error", errRest.Status())
	assert.EqualValues(t, "error while find one document", errRest.Message())
}
func TestErrNotFoundDeleteByID(t *testing.T) {
	elastic, err := New([]string{address})
	elastic.index = "posts"
	elastic.alias = elastic.index + "_alias"

	ps, errRest := NewPostStorage(*elastic)
	errRest = ps.DeleteByID(ctx, "710fd955-c2b8-4451-9ade-1cd8055d0dbe")

	assert.Nil(t, err)
	assert.NotNil(t, elastic)
	assert.NotNil(t, errRest)

	assert.EqualValues(t, http.StatusNotFound, errRest.Code())
	assert.EqualValues(t, "failed", errRest.Status())
	assert.EqualValues(t, "error document not found", errRest.Message())
}
func TestDeleteByID(t *testing.T) {
	elastic, err := New([]string{address})
	elastic.index = "posts"
	elastic.alias = elastic.index + "_alias"

	ps, errRest := NewPostStorage(*elastic)
	errRest = ps.Insert(ctx, newPost)
	errRest = ps.DeleteByID(ctx, newPost.ID)

	assert.Nil(t, err)
	assert.NotNil(t, elastic)
	assert.Nil(t, errRest)
}

// ===============================================

// ===============================================
func TestDocumentUpdateByID(t *testing.T) {
	elastic, err := New([]string{address})
	elastic.index = "posts"
	elastic.alias = elastic.index + "_alias"

	ps, errRest := NewPostStorage(*elastic)
	errRest = ps.Insert(ctx, newPost)
	errRest = ps.UpdateByID(ctx, updatedPost)

	assert.Nil(t, err)
	assert.Nil(t, errRest)
	assert.NotNil(t, ps)

	assert.EqualValues(t, id, updatedPost.ID)
	assert.EqualValues(t, "New Title from Update", updatedPost.Title)
	assert.EqualValues(t, "New Text from Update", updatedPost.Text)
	assert.EqualValues(t, []string{"tags2", "tags3"}, updatedPost.Tags)
}
func TestErrNotFoundUpdateByID(t *testing.T) {
	elastic, err := New([]string{address})
	elastic.index = "posts"
	elastic.alias = elastic.index + "_alias"

	ps, errRest := NewPostStorage(*elastic)
	errRest = ps.Insert(ctx, newPost)
	updatedPost.ID = "0c24230b-acfe-4b99-ac54-179e544158e0"
	errRest = ps.UpdateByID(ctx, updatedPost)

	assert.Nil(t, err)
	assert.NotNil(t, errRest)
	assert.NotNil(t, ps)

	assert.EqualValues(t, http.StatusNotFound, errRest.Code())
	assert.EqualValues(t, "failed", errRest.Status())
	assert.EqualValues(t, "error document not found", errRest.Message())
}
func TestErrUpdateByID(t *testing.T) {
	elastic, err := New([]string{wrongAddress})
	elastic.index = "posts"
	elastic.alias = elastic.index + "_alias"

	ps, errRest := NewPostStorage(*elastic)
	errRest = ps.Insert(ctx, newPost)
	updatedPost.ID = "0c24230b-acfe-4b99-ac54-179e544158e0"
	errRest = ps.UpdateByID(ctx, updatedPost)

	assert.Nil(t, err)
	assert.NotNil(t, errRest)
	assert.NotNil(t, ps)

	assert.EqualValues(t, http.StatusInternalServerError, errRest.Code())
	assert.EqualValues(t, "error", errRest.Status())
	assert.EqualValues(t, "error while updating new document", errRest.Message())
}
