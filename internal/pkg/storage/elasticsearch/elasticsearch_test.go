package elasticsearch

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	address       = "http://127.0.0.1:9206"
	wrongAddress  = "http://127.0.0.1:1234"
	id            = uuid.New().String()
	existingIndex = "test_posts"
	newIndex      = fmt.Sprintf("%s_%s", existingIndex, id)
)

func TestNewSuccessElasticConnection(t *testing.T) {
	elastic, err := New([]string{address})

	assert.Nil(t, err)
	assert.NotNil(t, elastic)
}

// Change to wrong address
func TestNonExistingIndex(t *testing.T) {
	elastic, err := New([]string{wrongAddress})
	errRest := elastic.CreateIndex(existingIndex)

	assert.Nil(t, err)
	assert.NotNil(t, errRest)
	assert.NotNil(t, elastic)

	assert.EqualValues(t, http.StatusInternalServerError, errRest.Code())
	assert.EqualValues(t, "error", errRest.Status())
	assert.EqualValues(t, "cannot check index existence", errRest.Message())
}

func TestExistingIndex(t *testing.T) {
	elastic, err := New([]string{address})
	errRest := elastic.CreateIndex(existingIndex)

	assert.Nil(t, err)
	assert.Nil(t, errRest)
	assert.NotNil(t, elastic)
}

func TestCreateNewIndex(t *testing.T) {
	elastic, err := New([]string{address})
	errRest := elastic.CreateIndex(newIndex)

	assert.Nil(t, err)
	assert.Nil(t, errRest)
	assert.NotNil(t, elastic)
}
