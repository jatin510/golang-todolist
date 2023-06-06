package handler

import (
	"log"

	"github.com/jatin510/golang-todolist/repository"
)

type Handler struct {
	TodoHandler TodoHandlerInterface
}

func InitHandler(repo repository.Repository, l *log.Logger) Handler {
	return Handler{
		TodoHandler: NewTodoHandler(repo, l),
	}
}
