package database

import (
	"context"
	"github.com/jackc/pgx/v5"
	"os"
)

type Postgres struct {
	db *pgx.Conn
}

func New() (*Postgres, error) {
	// ex: "postgresql://username:password@localhost:5432/mydb"
	db, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	return &Postgres{db: db}, nil
}

func (p *Postgres) Close() {
	p.db.Close(context.Background())
}
