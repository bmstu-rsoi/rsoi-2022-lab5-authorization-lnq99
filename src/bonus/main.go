package main

import (
	"database/sql"

	"bonus/config"
	"bonus/repository"
	"bonus/server"
	"bonus/service"

	"github.com/lnq99/rsoi-2022-lab3-fault-tolerance-lnq99/src/pkg/util"
)

func main() {
	var err error
	var db *sql.DB
	var cfg *config.Config

	if cfg, err = config.LoadConfig(); err != nil {
		panic(err)
	}

	if db, err = util.NewSqlDatabase(cfg.Db.Url); err != nil {
		panic(err)
	}
	defer db.Close()

	repo := repository.NewSqlRepository(db)

	svc := service.NewService(repo)

	svr := server.NewGinServer(svc, &cfg.Server)

	svr.Run()
}
