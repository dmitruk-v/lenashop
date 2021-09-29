package repository

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

var dbPool *pgxpool.Pool

func init() {
	var config *pgxpool.Config
	var err error

	config, err = pgxpool.ParseConfig("user=postgres password=419155 dbname=lenashop sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	dbPool, err = pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}
}
