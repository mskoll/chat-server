package repository

import (
	"chat-server/internal/model"
	"context"
	"github.com/jackc/pgx/v5"
)

type UserRepo struct {
	*pgx.Conn
}

func NewUser(db *pgx.Conn) *UserRepo {
	return &UserRepo{db}
}

func (r *UserRepo) Create(ctx context.Context, user model.User) (int, error) {

	query := `INSERT INTO users (username) VALUES ($1) RETURNING id`
	row := r.QueryRow(ctx, query, user.Username)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserRepo) Delete(ctx context.Context, id int) error {

	query := `UPDATE users SET deleted_at = now() WHERE id = $1`
	_, err := r.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepo) Chats(ctx context.Context, id int) (*[]model.ChatInfo, error) {

	query := `SELECT ct.id, ct.name, ct.created_at, ct.deleted_at FROM chats ct 
    INNER JOIN user_chat uct ON ct.id = uct.chat_id 
    INNER JOIN messages mt ON ct.id = mt.chat_id WHERE uct.user_id = $1 
                                                 ORDER BY mt.created_at`
	row := r.QueryRow(ctx, query, id)

	var chats []model.ChatInfo
	if err := row.Scan(&chats); err != nil {
		return nil, err
	}

	return &chats, nil
}
