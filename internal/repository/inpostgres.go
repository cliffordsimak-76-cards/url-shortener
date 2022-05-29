package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/model"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"log"
	"sync"
)

var createTableQuery = `
	create table if not exists urls (
	    correlation_id text unique,
	    user_id text not null,
		base_url text not null unique,
		url_id text not null
	);
	`

type InPostgres struct {
	db    *sql.DB
	mutex *sync.Mutex
}

func NewInPostgres(db *sql.DB) Repository {
	if _, err := db.Exec(createTableQuery); err != nil {
		log.Fatal(err)
	}
	return &InPostgres{
		db:    db,
		mutex: &sync.Mutex{},
	}
}

func (s *InPostgres) Create(
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

func (s *InPostgres) CreateBatch(urlModels []*model.URL) error {
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

func (s *InPostgres) Get(
	id string,
) (string, error) {
	var URL string
	err := s.db.QueryRow(
		"select base_url from urls where url_id=$1",
		id,
	).Scan(&URL)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", ErrNotFound
		}
		return "", err
	}
	return URL, nil
}

func (s *InPostgres) GetAll(
	userID string,
) ([]*model.URL, error) {
	rows, err := s.db.Query(
		"select base_url, url_id from urls where user_id=$1",
		userID,
	)
	if err != nil {
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
