package refreshtoken

import (
	"context"
	"github.com/evenyosua18/auth2/app/utils/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IRefreshTokenRepository interface {
	InsertRefreshToken(ctx context.Context, in interface{}) error
	GetRefreshToken(ctx context.Context, agg []bson.M) (interface{}, error)
	DeleteRefreshToken(ctx context.Context, filter bson.M) error
}

type RepositoryRefreshToken struct {
	col *mongo.Collection
	db  *db.MongoConnection
}

func NewRefreshTokenRepository(database *db.MongoConnection) IRefreshTokenRepository {
	collectionName := "refresh_tokens"

	return &RepositoryRefreshToken{
		col: database.DB.Collection(collectionName),
		db:  database,
	}
}
