package repository

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type TodoRepository struct {
	DB         *mongo.Database
	Collection *mongo.Collection
	l          *log.Logger
}

type TodoRepositoryInterface interface {
	GetTodo()
	GetAllTodo()
	SaveTodo()
}

var collection = "todos"

func NewTodoRepository(db *mongo.Database, l *log.Logger) TodoRepositoryInterface {
	return &TodoRepository{
		DB:         db,
		Collection: db.Collection(collection),
		l:          l,
	}
}

func (t *TodoRepository) GetTodo() {
	t.l.Println("get todo from repo")
}

func (t *TodoRepository) GetAllTodo() {
}

func (t *TodoRepository) SaveTodo() {

}
