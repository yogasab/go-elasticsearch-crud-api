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
