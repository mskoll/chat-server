package app

import (
	"chat-server/internal/database"
	"chat-server/internal/handler"
	"chat-server/internal/repository"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	db     *database.Postgres
	router *gin.Engine
	srv    *http.Server
}

func New() (app *App, err error) {

	app = &App{}

	app.db, err = database.New()
	if err != nil {
		return nil, err
	}

	repo := repository.New(app.db.Conn)
	handlers := handler.New(repo)

	app.router = handlers.Route()

	app.srv = &http.Server{
		Addr:    ":9000",
		Handler: app.router,
	}

	return app, err
}

func (a *App) Run() {

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := a.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()

	stop()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := a.srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}
	if err := a.db.Conn.Close(ctx); err != nil {
		log.Fatal("DB err: ", err)
	}

	log.Println("Server exiting")
}
