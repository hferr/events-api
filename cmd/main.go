package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hferr/events-api/config"
	"github.com/hferr/events-api/internal/app"
	"github.com/hferr/events-api/internal/httpjson"
	"github.com/hferr/events-api/internal/repositories"
	"github.com/hferr/events-api/internal/repositories/psql"
)

const fmtDbConnString = "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable"

func main() {
	cfg := config.New()

	db, err := initPostgresDb(cfg)
	if err != nil {
		log.Fatalf("could not initialize postgres: %v", err)
	}
	defer db.Close()

	repo := repositories.NewRepo(db)

	eventSvs := app.NewEventService(repo)

	h := httpjson.NewHandler(eventSvs)

	s := http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.ServerPort),
		Handler:      h.NewRouter(),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	fmt.Println("Starting server on port 8080")

	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}

func initPostgresDb(cfg *config.Cfg) (*sql.DB, error) {
	connString := fmt.Sprintf(
		fmtDbConnString,
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBName,
		cfg.DBPort,
	)

	postgres, err := psql.NewPostgres(connString)
	if err != nil {
		return nil, err
	}

	return postgres.Db, nil
}
