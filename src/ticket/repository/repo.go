package repository

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Repo Querier

func NewSqlRepository(db *sql.DB) Repo {
	return New(db)
}
