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
	Create(context.Context, *model.URL) error
	// CreateBatch.
	CreateBatch(context.Context, []*model.URL) error
	// UpdateBatch.
	UpdateBatch(context.Context, workers.DeleteTask) error
	// Get.
	Get(context.Context, string) (*model.URL, error)
	// GetAll.
	GetAll(context.Context, string) ([]*model.URL, error)
	// Stats.
	GetStats(context.Context) (*Stats, error)
}

// Errors.
var (
	// ErrNotFound.
	ErrNotFound = errors.New("URL is not found")
	// ErrAlreadyExists.
	ErrAlreadyExists = errors.New("URL already exists")
)
