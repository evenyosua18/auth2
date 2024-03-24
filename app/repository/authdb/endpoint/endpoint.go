package endpoint

import (
	"context"
	"github.com/evenyosua18/auth2/app/utils/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	IEndpointRepository interface {
		InsertEndpoint(ctx context.Context, in interface{}) error
		UpdateEndpoint(ctx context.Context, filter, in interface{}) error
		GetEndpoint(ctx context.Context, filter interface{}) (interface{}, error)
		GetEndpoints(ctx context.Context, filter interface{}) (interface{}, error)
	}

	RepositoryEndpoint struct {
		col *mongo.Collection
		db  *db.MongoConnection
	}
)

func NewEndpointRepository(database *db.MongoConnection) IEndpointRepository {
	collectionName := "endpoints"

	return &RepositoryEndpoint{
		col: database.DB.Collection(collectionName),
		db:  database,
	}
}
