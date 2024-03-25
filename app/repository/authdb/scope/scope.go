package scope

import (
	"context"
	"github.com/evenyosua18/auth2/app/utils/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type IScopeRepository interface {
	InsertScope(ctx context.Context, in interface{}) error
	GetScope(ctx context.Context, filter interface{}) (interface{}, error)
	//DeleteScope(ctx context.Context, filter interface{}) error
	UpdateScope(ctx context.Context, in, filter interface{}) error
	//ListScope(ctx context.Context, filter interface{}) (interface{}, error)
}

type RepositoryScope struct {
	col *mongo.Collection
	db  *db.MongoConnection
}

func NewScopeRepository(database *db.MongoConnection) IScopeRepository {
	collectionName := "scopes"

	return &RepositoryScope{
		col: database.DB.Collection(collectionName),
		db:  database,
	}
}
