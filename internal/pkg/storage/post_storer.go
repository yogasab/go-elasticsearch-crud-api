package storage

import (
	"context"
	"time"
)

type PostStore interface {
	Insert(ctx context.Context, post Post) error
}

type Post struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	Tags      []string   `json:"tags"`
	Text      string     `json:"text"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}
