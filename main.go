package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	l := log.Default()

	s := http.Server{
		Addr:         ":4000",
		ErrorLog:     l,
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
