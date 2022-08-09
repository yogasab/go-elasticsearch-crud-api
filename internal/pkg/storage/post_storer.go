package storage

import (
	"context"
	"time"
	"yogasab/go-elasticsearch-crud-api/internal/utils/http_errors"
)

type PostStore interface {
	Insert(ctx context.Context, post Post) http_errors.RestErrors
	FindByID(ctx context.Context, ID string) (*Post, http_errors.RestErrors)
	DeleteByID(ctx context.Context, ID string) http_errors.RestErrors
}

type Post struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	Tags      []string   `json:"tags"`
	Text      string     `json:"text"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}
