package db

import (
	"log"

	"github.com/jackc/pgx"
)

var Pool *pgx.ConnPool

func StartPool(connConfig pgx.ConnConfig) {
	var err error
	Pool, err = pgx.NewConnPool(extractConfig(connConfig))
	if err != nil {
		log.Fatalln("Unable to connect to database")
	}
}

func extractConfig(connConfig pgx.ConnConfig) pgx.ConnPoolConfig {
	var config pgx.ConnPoolConfig

	config.ConnConfig = connConfig

	return config
}
