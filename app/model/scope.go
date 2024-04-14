package model

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	ScopeModel struct {
		Id           primitive.ObjectID   `bson:"_id"`
		Service      string               `bson:"service"`
		Name         string               `json:"name"`
		Endpoints    []primitive.ObjectID `json:"endpoints"`
		ListEndpoint []EndpointModel      `json:"list_endpoint"`

		CreatedAt time.Time  `bson:"created_at"`
		UpdatedAt time.Time  `bson:"updated_at"`
		DeletedAt *time.Time `bson:"deleted_at"`
	}

	ScopeFilter struct {
		Pagination Pagination
		Sort       Sort
		Id         string
		Service    string
		Name       string
		Endpoints  []string

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
	if f.Id != "" {
		if id, err := primitive.ObjectIDFromHex(f.Id); err == nil {
			filter["_id"] = id
		}
	}

	// name
	if f.Name != "" {
		filter["name"] = primitive.Regex{Pattern: f.Name, Options: "i"}
	}

	// service
	if f.Service != "" {
		filter["service"] = f.Service
	}

	if len(f.Endpoints) != 0 {
		// convert to object id
		var endpoints []primitive.ObjectID
		for _, endpoint := range f.Endpoints {
			if id, err := primitive.ObjectIDFromHex(endpoint); err == nil {
				endpoints = append(endpoints, id)
			}
		}

		if len(endpoints) != 0 {
			filter["endpoints"] = bson.M{"$in": endpoints}
		}
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

	// always override endpoints
	updated["endpoints"] = m.Endpoints

	return bson.M{
		"$set": updated,
	}
}
