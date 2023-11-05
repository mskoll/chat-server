package repository

import (
	"chat-server/internal/model"
	"context"
	"github.com/jackc/pgx/v5"
)

type MessageRepo struct {
	*pgx.Conn
}

func NewMessage(db *pgx.Conn) *MessageRepo {
	return &MessageRepo{db}
}

func (r *MessageRepo) Create(ctx context.Context, message model.Message) (int, error) {

	query := `INSERT INTO messages (chat_id, user_id, content) VALUES ($1, $2, $3) RETURNING id`
	row := r.QueryRow(ctx, query, message.ChatId, message.UserId, message.Content)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
