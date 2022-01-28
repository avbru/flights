package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/avbru/flights/config"
	"log"
	"net/http"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfg)

	m, err := migrate.New("file://migrations", cfg.PgUrl)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Println("db is up to date")
		} else {
			log.Fatal(err)
		}
	}

	pool, err := pgxpool.Connect(context.Background(), cfg.PgUrl)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()

	srv := NewService(pool)

	http.HandleFunc("/flights", srv.flightsHandler)

	err = http.ListenAndServe(":"+cfg.Port, nil)
	if err != nil {
		log.Println(err.Error())
	}
}
