package oauthclient

import (
	"context"
	"github.com/evenyosua18/auth2/app/utils/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IOauthClientRepository interface {
	GetOauthClient(ctx context.Context, m bson.M) (interface{}, error)
	InsertOauthClient(ctx context.Context, in interface{}) error
}

type RepositoryOauthClient struct {
	col *mongo.Collection
	db  *db.MongoConnection
}

func NewOauthClientRepository(database *db.MongoConnection) IOauthClientRepository {
	collectionName := "oauth_clients"

	return &RepositoryOauthClient{
		col: database.DB.Collection(collectionName),
		db:  database,
	}
}
