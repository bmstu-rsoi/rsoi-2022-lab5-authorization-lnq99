package util

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewSqlDatabase(dbUrl string) (pool *sql.DB, err error) {
	maxTries := 10

	pool, err = sql.Open("pgx", dbUrl)
	if err != nil {
		return
	}

	for i := 0; i < maxTries; i++ {
		err = pool.Ping()
		if err == nil {
			break
		}
		time.Sleep(1 * time.Second)
		continue
	}

	return
}
