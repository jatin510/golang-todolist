package repository

import "go.mongodb.org/mongo-driver/mongo"

type Repository struct {
	TodoRepository TodoRepositoryInterface
}

func InitRepository(db *mongo.Database) Repository {
	return Repository{
		TodoRepository: InitTodoRepository(db),
	}
}
