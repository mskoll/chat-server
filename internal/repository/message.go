package repository

import (
	"chat-server/internal/model"
	"context"
	"github.com/jackc/pgx/v5"
)

type MessageRepo struct {
	db *pgx.Conn
}

func NewMessage(db *pgx.Conn) *MessageRepo {
	return &MessageRepo{db: db}
}

func (r *MessageRepo) Create(ctx context.Context, message model.Message) (int, error) {
	//TODO implement me
	panic("implement me")
}
