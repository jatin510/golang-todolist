package repository

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	TodoRepository TodoRepositoryInterface
}

func InitRepository(db *mongo.Database, l *log.Logger) Repository {
	return Repository{
		TodoRepository: NewTodoRepository(db, l),
	}
}
