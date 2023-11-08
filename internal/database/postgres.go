package database

import (
	"context"
	"github.com/jackc/pgx/v5"
	"os"
)

type Postgres struct {
	*pgx.Conn
}

func New() (*Postgres, error) {
	// ex: "postgresql://username:password@localhost:5432/mydb"
	db, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	return &Postgres{db}, nil
}
