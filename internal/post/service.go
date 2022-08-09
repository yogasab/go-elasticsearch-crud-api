package post

import (
	"context"
	"time"
	"yogasab/go-elasticsearch-crud-api/internal/pkg/storage"
	"yogasab/go-elasticsearch-crud-api/internal/utils/http_errors"

	"github.com/google/uuid"
)

type PostService interface {
	InsertDocument(ctx context.Context, request InsertDocumentRequest) (*InsertDocumentRequest, http_errors.RestErrors)
	FindDocumentByID(ctx context.Context, ID string) (*storage.Post, http_errors.RestErrors)
	DeleteDocumentByID(ctx context.Context, ID string) (bool, http_errors.RestErrors)
	UpdateDocumentByID(ctx context.Context, request UpdateDocumentRequest) (bool, http_errors.RestErrors)
}

type postService struct {
	storage storage.PostStore
}

func NewPostService(storage storage.PostStore) PostService {
	return &postService{storage: storage}
}

func (s postService) InsertDocument(ctx context.Context, request InsertDocumentRequest) (*InsertDocumentRequest, http_errors.RestErrors) {
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
		return nil, err
	}
	return &request, nil
}

func (s postService) FindDocumentByID(ctx context.Context, ID string) (*storage.Post, http_errors.RestErrors) {
	post, err := s.storage.FindByID(ctx, ID)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s postService) DeleteDocumentByID(ctx context.Context, ID string) (bool, http_errors.RestErrors) {
	isDeleted := true
	if err := s.storage.DeleteByID(ctx, ID); err != nil {
		isDeleted = false
		return isDeleted, err
	}
	return isDeleted, nil
}

func (s postService) UpdateDocumentByID(ctx context.Context, request UpdateDocumentRequest) (bool, http_errors.RestErrors) {
	post := storage.Post{}
	post.ID = request.ID
	post.Title = request.Title
	post.Text = request.Text
	post.Tags = request.Tags

	isUpdated := true
	err := s.storage.UpdateByID(ctx, post)
	if err != nil {
		isUpdated = false
		return isUpdated, err
	}
	return isUpdated, nil
}
