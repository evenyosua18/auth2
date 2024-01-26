package model

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	AccessTokenModel struct {
		Id        primitive.ObjectID `bson:"_id"`
		ExpiredAt time.Time          `bson:"expired_at"`
		UserId    primitive.ObjectID `bson:"user_id"`

		CreatedAt time.Time  `bson:"created_at"`
		UpdatedAt time.Time  `bson:"updated_at"`
		DeletedAt *time.Time `bson:"deleted_at"`
	}

	AccessTokenFilter struct {
		Id        *primitive.ObjectID
		UserId    *primitive.ObjectID
		ExpiredAt *time.Time

		Deleted bool
	}
)

func (f AccessTokenFilter) Filter() bson.M {
	filter := bson.M{}

	// deleted
	if !f.Deleted {
		filter["deleted_at"] = nil
	}

	// id
	if f.Id != nil {
		filter["_id"] = f.Id
	}

	// user id
	if f.UserId != nil {
		filter["user_id"] = f.UserId
	}

	// expired at
	if f.ExpiredAt != nil {
		filter["expired_at"] = bson.M{
			"$gte": f.ExpiredAt,
		}
	}

	return filter
}

func (m AccessTokenModel) SetCreated() AccessTokenModel {
	if m.Id.IsZero() {
		m.Id = primitive.NewObjectID()
	}

	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()

	return m
}

func (m AccessTokenModel) Update() bson.M {
	// updated at
	updated := bson.M{
		"updated_at": time.Now(),
	}

	if !m.ExpiredAt.IsZero() {
		updated["expired_at"] = m.ExpiredAt
	}

	return bson.M{
		"$set": updated,
	}
}
