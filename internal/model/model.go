package model

import "time"

type User struct {
	Username string `json:"username" validate:"required"`
}

type UserInfo struct {
	Id        int
	Username  string
	CreatedAt *time.Time `db:"created_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

type Chat struct {
	Name  string `json:"name" validate:"required"`
	Users []int  `json:"users" validate:"required"`
}

type ChatInfo struct {
	Id        int
	Name      string
	CreatedAt *time.Time `db:"created_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

type Message struct {
	ChatId  int    `json:"chat" validate:"required"`
	UserId  int    `json:"author" validate:"required"`
	Content string `json:"text" validate:"required"`
}

type MessageInfo struct {
	Id        int        `json:"id"`
	ChatId    int        `json:"chat"`
	UserId    int        `json:"author"`
	Content   string     `json:"text"`
	CreatedAt *time.Time `db:"created_at"`
}
