package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type OauthClientModel struct {
	Id           primitive.ObjectID `bson:"_id"`
	ClientId     string             `bson:"client_id"`
	ClientSecret string             `bson:"client_secret"`
	Uri          string             `bson:"uri"`
	Scopes       string             `bson:"scopes"`
	ClientType   string             `bson:"client_type"`

	CreatedAt time.Time  `bson:"created_at"`
	UpdatedAt time.Time  `bson:"updated_at"`
	DeletedAt *time.Time `bson:"deleted_at"`
}
