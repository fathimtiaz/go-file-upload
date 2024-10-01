package api

import (
	"go-file-upload/config"
	"go-file-upload/internal/repository/postgres"
	"go-file-upload/internal/service"
	"log"

	_ "github.com/lib/pq"
)

func Run() {
	cfg := config.LoadDefault()

	postgresDB, err := postgres.NewSqlDB("postgres", cfg.DB.ConnStr.String())
	if err != nil {
		log.Fatal(err)
	}

	fileSvc := service.NewFileSvc(cfg, postgresDB)

	Router(cfg, fileSvc).Run(":" + cfg.App.Port.String())
}
