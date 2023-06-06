package route

import (
	"net/http"

	"github.com/jatin510/golang-todolist/handler"
)

func InitRoute(handler handler.Handler) http.Handler {
	sm := http.NewServeMux()

	sm.HandleFunc("/todo/get", handler.TodoHandler.GetTodo)
	sm.HandleFunc("/todo/create", handler.TodoHandler.SaveTodo)
	return sm

}
