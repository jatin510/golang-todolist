package repository

import "go.mongodb.org/mongo-driver/mongo"

type TodoRepository struct {
	DB         *mongo.Database
	Collection *mongo.Collection
}

type TodoRepositoryInterface interface {
	getTodo()
	getAllTodo()
	saveTodo()
}

var collection = "todos"

func InitTodoRepository(db *mongo.Database) TodoRepositoryInterface {
	return &TodoRepository{
		DB: db,
		// Collection: db.Collection(collection),
	}
}

func (t *TodoRepository) getTodo() {

}

func (t *TodoRepository) getAllTodo() {

}

func (t *TodoRepository) saveTodo() {

}
