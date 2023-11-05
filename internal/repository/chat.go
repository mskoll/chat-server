package repository

import (
	"chat-server/internal/model"
	"context"
	"github.com/jackc/pgx/v5"
)

type ChatRepo struct {
	*pgx.Conn
}

func NewChat(db *pgx.Conn) *ChatRepo {
	return &ChatRepo{db}
}

func (r *ChatRepo) Create(ctx context.Context, chat model.Chat) (int, error) {

	tx, err := r.Begin(ctx)
	if err != nil {
		return 0, err
	}

	query := `INSERT INTO chats (name) VALUES ($1) RETURNING id`
	row := tx.QueryRow(ctx, query, chat.Name)

	var id int
	if err = row.Scan(&id); err != nil {
		tx.Rollback(ctx)
		return 0, err
	}

	query = `INSERT INTO user_chat (user_id, chat_id) VALUES ($1, $2)`
	for _, user := range chat.Users {
		_, err = tx.Exec(ctx, query, user, id)
		if err != nil {
			tx.Rollback(ctx)
			return 0, err
		}
	}

	return id, tx.Commit(ctx)
}

func (r *ChatRepo) Delete(ctx context.Context, id int) error {

	query := `UPDATE chats SET deleted_at = now() WHERE id = $1`
	_, err := r.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *ChatRepo) Users(ctx context.Context, id int) (*[]model.UserInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ChatRepo) Messages(ctx context.Context, id int) (*[]model.MessageInfo, error) {
	//TODO implement me
	panic("implement me")
}
