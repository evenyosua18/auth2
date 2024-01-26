package user

import (
	"context"
	"github.com/evenyosua18/auth2/app/utils/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserRepository interface {
	GetUser(ctx context.Context, filter interface{}) (interface{}, error)
	InsertUser(ctx context.Context, in interface{}) error
}

type RepositoryUser struct {
	col *mongo.Collection
	db  *db.MongoConnection
}

func NewUserRepository(database *db.MongoConnection) IUserRepository {
	collectionName := "users"

	return &RepositoryUser{
		col: database.DB.Collection(collectionName),
		db:  database,
	}
}
