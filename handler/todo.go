package handler

import (
	"log"
	"net/http"

	"github.com/jatin510/golang-todolist/repository"
)

type TodoHandler struct {
	repository repository.Repository
	l          *log.Logger
}

type TodoHandlerInterface interface {
	GetTodo(w http.ResponseWriter, r *http.Request)
	SaveTodo(w http.ResponseWriter, r *http.Request)
}

func NewTodoHandler(repository repository.Repository, l *log.Logger) TodoHandlerInterface {
	return &TodoHandler{
		repository: repository,
	}
}

func (t *TodoHandler) GetTodo(w http.ResponseWriter, r *http.Request) {
	t.l.Println("get todo")
}

func (t *TodoHandler) SaveTodo(w http.ResponseWriter, r *http.Request) {
	t.l.Println("save todo")
}