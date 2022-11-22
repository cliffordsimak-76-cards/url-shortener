package workers

import (
	"context"
	"sync"
)

// DeleteTask.
type DeleteTask struct {
	UserID string
	UrlsID []string
}

// DeleteWorker.
type DeleteWorker struct {
	repo  Repository
	Tasks <-chan DeleteTask
}

// New.
func New(
	repo Repository,
	tasks <-chan DeleteTask,
) *DeleteWorker {
	return &DeleteWorker{
		repo:  repo,
		Tasks: tasks,
	}
}

// Run.
func (w *DeleteWorker) Run(ctx context.Context) {
	wg := sync.WaitGroup{}
	for {
		select {
		case <-ctx.Done():
			wg.Wait()
			return
		case task := <-w.Tasks:
			wg.Add(1)
			go func() {
				defer wg.Done()
				w.AddWorker(ctx, task)
			}()
		}
	}
}

// Repository.
type Repository interface {
	UpdateBatch(ctx context.Context, task DeleteTask) error
}

// AddWorker.
func (w *DeleteWorker) AddWorker(ctx context.Context, task DeleteTask) {
	w.repo.UpdateBatch(ctx, task)
}
