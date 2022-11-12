package workers

import (
	"context"
	"sync"
)

type DeleteTask struct {
	UserID string
	UrlsID []string
}

type DeleteWorker struct {
	repo  Repository
	Tasks <-chan DeleteTask
}

func New(
	repo Repository,
	tasks <-chan DeleteTask,
) *DeleteWorker {
	return &DeleteWorker{
		repo:  repo,
		Tasks: tasks,
	}
}

// run.
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

type Repository interface {
	UpdateBatch(ctx context.Context, task DeleteTask) error
}

func (w *DeleteWorker) AddWorker(ctx context.Context, task DeleteTask) {
	w.repo.UpdateBatch(ctx, task)
}
