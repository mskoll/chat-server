package repository

import (
	"chat-server/internal/model"
	"context"
	"github.com/jackc/pgx/v5"
)

type ChatRepo struct {
	db *pgx.Conn
}

func NewChat(db *pgx.Conn) *ChatRepo {
	return &ChatRepo{db: db}
}

func (r *ChatRepo) Create(ctx context.Context, chat model.Chat) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ChatRepo) Delete(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func (r *ChatRepo) Users(ctx context.Context, id int) (*[]model.UserInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ChatRepo) Messages(ctx context.Context, id int) (*[]model.MessageInfo, error) {
	//TODO implement me
	panic("implement me")
}
