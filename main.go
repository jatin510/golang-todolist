package main

import (
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

	s := http.Server{
		Addr:         ":4000",
		ErrorLog:     l,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		l.Println("starting server on port 4000")

		err := s.ListenAndServe()
		if err != nil {
			l.Println("server stopped ", err)
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-done
	l.Print("Server Stopped")
}
