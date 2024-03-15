package model

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	EndpointModel struct {
		Id         primitive.ObjectID `bson:"_id"`
		Service    string             `bson:"service"`
		Endpoint   string             `bson:"endpoint"`
		IsGenerate *bool              `bson:"is_generate"`
		StillExist *bool              `bson:"still_exist"` // still exist in last sync process

		CreatedAt time.Time  `bson:"created_at"`
		UpdatedAt time.Time  `bson:"updated_at"`
		DeletedAt *time.Time `bson:"deleted_at"`
	}

	EndpointFilter struct {
		Id         *primitive.ObjectID
		Service    string
		Endpoint   string
		IsGenerate *bool

		Deleted bool
	}
)

func (m EndpointModel) SetCreated() EndpointModel {
	if m.Id.IsZero() {
		m.Id = primitive.NewObjectID()
	}

	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()

	return m
}

func (f EndpointFilter) Filter() bson.M {
	filter := bson.M{}

	// deleted at
	if !f.Deleted {
		filter["deleted_at"] = nil
	}

	// id
	if f.Id != nil {
		filter["_id"] = f.Id
	}

	// service
	if f.Service != "" {
		filter["service"] = f.Service
	}

	// endpoint
	if f.Endpoint != "" {
		filter["endpoint"] = primitive.Regex{Pattern: f.Endpoint, Options: "i"}
	}

	// is generate
	if f.IsGenerate != nil {
		filter["is_generate"] = true
	}

	return filter
}

func (m EndpointModel) Update() bson.M {
	// updated at
	updated := bson.M{
		"updated_at": time.Now(),
	}

	// is_generate
	if m.IsGenerate != nil {
		updated["is_generate"] = m.IsGenerate
	}

	// service
	if m.Service != "" {
		updated["service"] = m.Service
	}

	// endpoint
	if m.Endpoint != "" {
		updated["endpoint"] = m.Endpoint
	}

	return bson.M{
		"$set": updated,
	}
}
