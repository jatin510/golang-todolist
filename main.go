package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jatin510/golang-todolist/database"
	"github.com/jatin510/golang-todolist/handler"
	"github.com/jatin510/golang-todolist/repository"
	"github.com/jatin510/golang-todolist/route"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	DB        *mongo.Database
	DB_Client *mongo.Client
)

func init() {
	DB_Client, DB = database.InitDatabase()
}

func main() {

	l := log.Default()

	// init repository
	repo := repository.InitRepository(DB, l)

	// init handler
	handler := handler.InitHandler(repo, l)

	// init router
	router := route.InitRoute(handler)

	port := "4000"

	s := http.Server{
		Addr:         ":" + port,
		ErrorLog:     l,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Print("Server Started on port ", port)
	<-done
	log.Print("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// extra handling here
		cancel()
	}()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Print("Server Exited Properly")
}
