package model

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	OauthClientModel struct {
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

	OauthClientFilter struct {
		Id       *primitive.ObjectID
		ClientId string

		Deleted bool
	}
)

func (m OauthClientModel) SetCreated() OauthClientModel {
	if m.Id.IsZero() {
		m.Id = primitive.NewObjectID()
	}

	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()

	return m
}

func (f OauthClientFilter) Filter() bson.M {
	filter := bson.M{}

	// deleted at
	if !f.Deleted {
		filter["deleted_at"] = nil
	}

	// id
	if f.Id != nil && !f.Id.IsZero() {
		filter["_id"] = f.Id
	}

	// client id
	if f.ClientId != "" {
		filter["client_id"] = f.ClientId
	}

	return filter
}
