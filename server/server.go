package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Run() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	g := router.Group("/api")
	AuditRoutes(g)

	server := &http.Server{
		Addr:    viper.GetString("host"),
		Handler: router,
	}

	// this code is ran in a separate thread
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("listen err:", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// this code is blocked until the quit channel receives a message, which will
	// only happen if a SIGINT or SIGTERM signal is given to the process
	gracefulShutdown(server, <-quit)
}

func gracefulShutdown(server *http.Server, signal os.Signal) {
	log.Printf("Received %v, server shutting down...", signal)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("shutdown err:", err)
	}

	log.Println("Goodbye!")
}
