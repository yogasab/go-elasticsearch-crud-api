package post

import (
	"context"
	"errors"
	"fmt"
	"time"
	"yogasab/go-elasticsearch-crud-api/internal/pkg/storage"

	"github.com/google/uuid"
)

type PostService interface {
	InsertDocument(ctx context.Context, request InsertDocumentRequest) (*InsertDocumentRequest, error)
}

type postService struct {
	storage storage.PostStore
}

func NewPostService(storage storage.PostStore) PostService {
	return &postService{storage: storage}
}

func (s postService) InsertDocument(ctx context.Context, request InsertDocumentRequest) (*InsertDocumentRequest, error) {
	id := uuid.New().String()
	created_at := time.Now().UTC()

	post := storage.Post{}
	post.ID = id
	post.Title = request.Title
	post.Text = request.Text
	post.Tags = request.Tags
	post.CreatedAt = &created_at

	err := s.storage.Insert(ctx, post)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error while create new document service %v", err))
	}
	return &request, nil
}
