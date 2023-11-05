package repository

import (
	"chat-server/internal/model"
	"context"
	"github.com/jackc/pgx/v5"
)

type User interface {
	Create(ctx context.Context, user model.User) (int, error)
	Delete(ctx context.Context, id int) error
	Chats(ctx context.Context, id int) (*[]model.ChatInfo, error)
	//сортировка по времени создания последнего сообщения (от позднего к раннему)
}

type Chat interface {
	Create(ctx context.Context, chat model.Chat) (int, error)
	Delete(ctx context.Context, id int) error
	Users(ctx context.Context, id int) (*[]model.UserInfo, error)
	Messages(ctx context.Context, id int) (*[]model.MessageInfo, error)
	//сортировка по времени создания сооб (от раннего к позднему)
}

type Message interface {
	Create(ctx context.Context, message model.Message) (int, error)
}

type Repository struct {
	User
	Chat
	Message
}

func New(db *pgx.Conn) *Repository {
	return &Repository{
		User:    NewUser(db),
		Chat:    NewChat(db),
		Message: NewMessage(db),
	}
}
