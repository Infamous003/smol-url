package app

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type Application struct {
	httpServer *http.Server
	router     *gin.Engine
	// db *gorm.DB // add this later
}

func New() *Application {
	// Load up conigs

	// Connect to db

	// create the Server struct and return

	// Create gin router
	r := gin.Default()

	return &Application{
		router: r,
	}
}

func (app *Application) Start(addr string) error {
	app.httpServer = &http.Server{
		Addr:    addr,
		Handler: app.router,
	}

	// quit channel captures signals like ctrl+c
	quit := make(chan os.Signal, 1)

	// shutdownErr stores err returned from .Shutdwon()
	shutdownErr := make(chan error, 1)

	go func() {
		// Notify listens to syscalls and sends them to quit channel
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		// Explicitly blocking, since Notify is non blocking
		<-quit

		log.Println("Shutdown signal received. Shutting down the server...")
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		// need a way to capture .Shutdown() err, which can be nil or non nil
		shutdownErr <- app.httpServer.Shutdown(ctx)
	}()

	log.Printf("Server starting on %s", addr)
	err := app.httpServer.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return <-shutdownErr
}

func (app *Application) Router() *gin.Engine {
	return app.router
}
