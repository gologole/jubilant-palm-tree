package repository

import (
	"cmd/main.go/config"
	"context"
	"github.com/jackc/pgx"
	"log"
	"time"
)

type Database interface {
}
type database struct {
	conn   *pgx.Conn
	config *config.Config
}

func NewDatabase(cfg *config.Config) (Database, error) {
	conn, err := pgx.Connect(pgx.ConnConfig{
		Host:     cfg.Postgres.Host,
		Port:     cfg.Postgres.Port,
		Database: cfg.Postgres.Database,
		User:     cfg.Postgres.User,
		Password: cfg.Postgres.Password,
	})
	if err != nil {
		return nil, err
	}
	for {
		err := conn.Ping(context.Background())
		if err == nil {
			break
		}
		log.Println("Database is still starting up, retrying in 2 seconds...")
		time.Sleep(2 * time.Second)
	}

	return &database{conn: conn, config: cfg}, nil
}
