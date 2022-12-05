package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/workers"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/model"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/labstack/gommon/log"
)

// CreateTableQuery.
var CreateTableQuery = `
	create table if not exists urls (
	    correlation_id text unique,
	    user_id text not null,
		base_url text not null unique,
		url_id text not null,
		deleted bool default false
	);
	`

// slq db.
type InDatabase struct {
	db *sql.DB
}

// NewInDatabase.
func NewInDatabase(db *sql.DB) Repository {
	log.Info("start database repo")
	return &InDatabase{
		db: db,
	}
}

// Create.
func (s *InDatabase) Create(
	ctx context.Context,
	urlModel *model.URL,
) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(
		ctx,
		"INSERT INTO urls(correlation_id, user_id, base_url, url_id) VALUES($1,$2,$3,$4)",
	)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.ExecContext(
		ctx,
		urlModel.CorrelationID,
		urlModel.UserID,
		urlModel.Original,
		urlModel.Short,
	); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case pgerrcode.UniqueViolation:
				return ErrAlreadyExists
			}
		}
		return err
	}

	return tx.Commit()
}

// CreateBatch.
func (s *InDatabase) CreateBatch(
	ctx context.Context,
	urlModels []*model.URL,
) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(
		ctx,
		"INSERT INTO urls(correlation_id, user_id, base_url, url_id) VALUES($1,$2,$3,$4)",
	)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, urlModel := range urlModels {
		if _, err = stmt.ExecContext(
			ctx,
			urlModel.CorrelationID,
			urlModel.UserID,
			urlModel.Original,
			urlModel.Short,
		); err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				switch pgErr.Code {
				case pgerrcode.UniqueViolation:
					return ErrAlreadyExists
				}
			}
			return err
		}
	}

	return tx.Commit()
}

// Get.
func (s *InDatabase) Get(
	ctx context.Context,
	id string,
) (*model.URL, error) {
	URL := &model.URL{}
	err := s.db.QueryRowContext(ctx,
		"select base_url, url_id, deleted from urls where url_id=$1",
		id,
	).Scan(&URL.Original, &URL.Short, &URL.Deleted)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return URL, nil
}

// GetAll.
func (s *InDatabase) GetAll(
	ctx context.Context,
	userID string,
) ([]*model.URL, error) {
	rows, err := s.db.QueryContext(ctx,
		"select base_url, url_id from urls where user_id=$1",
		userID,
	)
	if rows.Err() != nil {
		return nil, err
	}
	defer rows.Close()

	var urls []*model.URL

	for rows.Next() {
		url := &model.URL{}
		err = rows.Scan(&url.Original, &url.Short)
		if err != nil {
			return nil, err
		}
		urls = append(urls, url)
	}

	return urls, nil
}

// UpdateBatch.
func (s *InDatabase) UpdateBatch(
	ctx context.Context,
	task workers.DeleteTask,
) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(
		ctx,
		"UPDATE urls SET deleted = true WHERE url_id = any($1) AND user_id = $2",
	)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.ExecContext(
		ctx,
		task.UrlsID,
		task.UserID,
	); err != nil {
		return err
	}

	return tx.Commit()
}

// Stats.
func (s *InDatabase) GetStats(
	ctx context.Context,
) (*Stats, error) {
	stats := &Stats{}

	err := s.db.QueryRowContext(ctx, "SELECT COUNT(1), COUNT(DISTINCT user_id) FROM urls").
		Scan(stats.LinksCount, stats.UsersCount)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return stats, nil
}
