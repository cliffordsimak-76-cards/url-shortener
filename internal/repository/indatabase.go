package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/model"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/labstack/gommon/log"
)

var CreateTableQuery = `
	create table if not exists urls (
	    correlation_id text unique,
	    user_id text not null,
		base_url text not null unique,
		url_id text not null
	);
	`

type InDatabase struct {
	db *sql.DB
}

func NewInDatabase(db *sql.DB) Repository {
	log.Info("start database repo")
	return &InDatabase{
		db: db,
	}
}

func (s *InDatabase) Create(
	urlModel *model.URL,
) error {
	ctx := context.Background()
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

func (s *InDatabase) CreateBatch(urlModels []*model.URL) error {
	ctx := context.Background()
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

func (s *InDatabase) Get(id string) (*model.URL, error) {
	URL := &model.URL{}
	err := s.db.QueryRow(
		"select base_url, url_id from urls where url_id=$1",
		id,
	).Scan(&URL.Original, &URL.Short)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return URL, nil
}

func (s *InDatabase) GetAll(userID string) ([]*model.URL, error) {
	rows, err := s.db.Query(
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
