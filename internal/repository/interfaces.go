package repository

import (
	"context"
	"errors"

	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/workers"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/model"
)

// Repo interfaces.
type Repository interface {
	// Create.
	Create(urlModel *model.URL) error
	// CreateBatch.
	CreateBatch(urlModels []*model.URL) error
	// UpdateBatch.
	UpdateBatch(ctx context.Context, task workers.DeleteTask) error
	// Get.
	Get(id string) (*model.URL, error)
	// GetAll.
	GetAll(userID string) ([]*model.URL, error)
}

var (
	// ErrNotFound.
	ErrNotFound = errors.New("URL is not found")
	// ErrAlreadyExists.
	ErrAlreadyExists = errors.New("URL already exists")
)
