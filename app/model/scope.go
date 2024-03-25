package model

import (
	"go.mongodb.org/mongo-driver/bson"
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

func (f ScopeFilter) Filter() bson.M {
	filter := bson.M{}

	// deleted at
	if !f.Deleted {
		filter["deleted_at"] = nil
	}

	// id
	if f.Id != nil && !f.Id.IsZero() {
		filter["_id"] = f.Id
	}

	// name
	if f.Name != "" {
		filter["name"] = primitive.Regex{Pattern: f.Name, Options: "i"}
	}

	// service
	if f.Service != "" {
		filter["service"] = f.Service
	}

	// TODO make sure again it works
	if len(f.Endpoints) != 0 {
		filter["endpoints"] = bson.M{"$in": f.Endpoints}
	}

	return filter
}

func (m ScopeModel) Update() bson.M {
	// updated at
	updated := bson.M{
		"updated_at": time.Now(),
	}

	// name
	if m.Name != "" {
		updated["name"] = m.Name
	}

	//service
	if m.Service != "" {
		updated["service"] = m.Service
	}

	// TODO add & delete endpoint

	return bson.M{
		"$set": updated,
	}
}
