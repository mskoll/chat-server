package handler

import (
	"chat-server/internal/repository"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	*repository.Repository
}

func New(repository *repository.Repository) *Handler {
	return &Handler{repository}
}

func (h *Handler) Route() *gin.Engine {
	router := gin.New()

	users := router.Group("/users")
	users.POST("/add", h.createUser)
	users.GET("/chats/:id", h.userChats)
	users.DELETE("/delete/:id", h.deleteUser)

	chats := router.Group("/chats")
	chats.POST("/add", h.createChat)
	chats.GET("/users/:id", h.chatUsers)
	chats.GET("/messages/:id", h.chatMessages)
	chats.DELETE("/:id", h.deleteChat)

	messages := router.Group("/messages")
	messages.POST("/add", h.createMessage)

	return router
}
