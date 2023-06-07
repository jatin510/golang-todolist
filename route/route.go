package route

import (
	"net/http"

	"github.com/jatin510/golang-todolist/handler"
)

func InitRoute(handler handler.Handler) http.Handler {
	sm := http.NewServeMux()

	sm.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handler.TodoHandler.GetTodo(w, r)
		}

		if r.Method == http.MethodPost {
			handler.TodoHandler.SaveTodo(w, r)
		}
	})
	return sm

}
