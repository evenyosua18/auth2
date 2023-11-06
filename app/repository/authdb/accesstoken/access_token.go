package accesstoken

import (
	"context"
	"github.com/evenyosua18/auth2/app/utils/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IAccessTokenRepository interface {
	InsertAccessToken(ctx context.Context, in interface{}) error
	GetAccessToken(ctx context.Context, filter bson.M) (interface{}, error)
	DeleteAccessToken(ctx context.Context, filter bson.M) error
	UpdateAccessToken(ctx context.Context, filter, in bson.M) (interface{}, error)
}

type RepositoryAccessToken struct {
	col *mongo.Collection
	db  *db.MongoConnection
}

func NewAccessTokenRepository(database *db.MongoConnection) IAccessTokenRepository {
	collectionName := "access_tokens"

	return &RepositoryAccessToken{
		col: database.DB.Collection(collectionName),
		db:  database,
	}
}
