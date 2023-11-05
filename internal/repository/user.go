package repository

import (
	"chat-server/internal/model"
	"context"
	"github.com/jackc/pgx/v5"
)

type UserRepo struct {
	db *pgx.Conn
}

func NewUser(db *pgx.Conn) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(ctx context.Context, user model.User) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (r *UserRepo) Delete(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func (r *UserRepo) Chats(ctx context.Context, id int) (*[]model.ChatInfo, error) {
	//TODO implement me
	panic("implement me")
}
