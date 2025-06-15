package helper

import (
	"database/sql"
	"kido1611/notes-backend-go/internal/db/sqlc"

	"github.com/sirupsen/logrus"
)

func DbTransaction[T any](db *sql.DB, log *logrus.Logger, tasks func(*sqlc.Queries) (T, error)) (T, error) {
	var zero T
	tx, err := db.Begin()
	if err != nil {
		log.Warnf("Failed starting database transaction: %+v", err)
		return zero, err
	}
	defer tx.Rollback()

	qtx := sqlc.New(tx)

	result, err := tasks(qtx)
	if err != nil {
		log.Warnf("Failed executing tasks: %+v", err)
		return zero, err
	}

	err = tx.Commit()
	if err != nil {
		log.Warnf("Failed commit database: %+v", err)
		return zero, err
	}

	return result, nil
}
