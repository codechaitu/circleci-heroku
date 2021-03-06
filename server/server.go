package server

import (
	"context"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"fmt"

	"github.com/gin-gonic/gin"
)

// CreateRouter creates and configures a server
func CreateRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.Use(DB(db))
	setupRoutes(router)
	return router
}

// StartServer starts given server, supporting graceful shutdown of the server
func StartServer(router *gin.Engine) {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	fmt.Println(port)
	srv := &http.Server{
		Addr:    ":"+port,
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
