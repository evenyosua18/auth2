package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	ScopeModel struct {
		Id        primitive.ObjectID   `bson:"_id"`
		Service   string               `bson:"service"`
		Name      string               `json:"name"`
		Endpoints []primitive.ObjectID `json:"endpoints"`

		CreatedAt time.Time  `bson:"created_at"`
		UpdatedAt time.Time  `bson:"updated_at"`
		DeletedAt *time.Time `bson:"deleted_at"`
	}

	ScopeFilter struct {
		Id        *primitive.ObjectID
		Service   string
		Name      string
		Endpoints []primitive.ObjectID

		Deleted bool
	}
)

func (m ScopeModel) SetCreated() ScopeModel {
	if m.Id.IsZero() {
		m.Id = primitive.NewObjectID()
	}

	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()

	return m
}
