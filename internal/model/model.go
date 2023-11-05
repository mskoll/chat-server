package model

import "time"

type User struct {
	Username string `json:"username" validate:"required"`
}

type UserInfo struct {
	Id        int        `json:"id"`
	Username  string     `json:"username"`
	CreatedAt *time.Time `json:"createdAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
}

type Chat struct {
	Name  string `json:"name" validate:"required"`
	Users []int  `json:"users" validate:"required"`
}

type ChatInfo struct {
	Id        int        `json:"id"`
	Name      string     `json:"name"`
	Users     []UserInfo `json:"users"`
	CreatedAt *time.Time `json:"createdAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
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
	CreatedAt *time.Time `json:"createdAt"`
}
