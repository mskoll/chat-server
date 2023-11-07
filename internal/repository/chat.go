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

	query := `SELECT ut.id, ut.username, ut.created_at, ut.deleted_at FROM users ut 
    INNER JOIN user_chat uct ON ut.id = uct.user_id WHERE uct.chat_id = $1 AND uct.deleted_at IS NULL`
	row := r.QueryRow(ctx, query, id)

	var users []model.UserInfo
	if err := row.Scan(&users); err != nil {
		return nil, err
	}

	return &users, nil
}

func (r *ChatRepo) Messages(ctx context.Context, id int) (*[]model.MessageInfo, error) {

	query := `SELECT * FROM messages WHERE chat_id = $1 ORDER BY created_at`
	row := r.QueryRow(ctx, query, id)

	var messages []model.MessageInfo
	if err := row.Scan(&messages); err != nil {
		return nil, err
	}

	return &messages, nil
}
