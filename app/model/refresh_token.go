package model

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	RefreshTokenModel struct {
		Id            primitive.ObjectID `bson:"_id"`
		AccessTokenId primitive.ObjectID `bson:"access_token_id"`
		UserId        primitive.ObjectID `bson:"user_id"`
		RefreshToken  string             `bson:"refresh_token"`
		Count         int                `bson:"count"`

		AccessToken *AccessTokenModel
		User        *UserModel

		CreatedAt time.Time  `bson:"created_at"`
		UpdatedAt time.Time  `bson:"updated_at"`
		DeletedAt *time.Time `bson:"deleted_at"`
	}

	RefreshTokenFilter struct {
		Id           *primitive.ObjectID
		Deleted      bool
		RefreshToken string
	}
)

func (m RefreshTokenModel) SetCreated() {
	if m.Id.IsZero() {
		m.Id = primitive.NewObjectID()
	}

	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
}

func (f RefreshTokenFilter) Aggregate(additions ...bson.M) []bson.M {
	var res []bson.M

	// match
	res = append(res, bson.M{
		"$match": f.Filter(),
	})

	// add additions
	res = append(res, additions...)

	// lookup access token
	res = append(res, bson.M{
		"$lookup": bson.M{
			"from":         "access_tokens",
			"localField":   "access_token_id",
			"foreignField": "_id",
			"as":           "access_token",
		},
	})

	// unwind access token
	res = append(res, bson.M{
		"$unwind": bson.M{
			"path":                       "$access_token",
			"preserveNullAndEmptyArrays": true,
		},
	})

	return res
}

func (f RefreshTokenFilter) Filter() bson.M {
	filter := bson.M{}

	// deleted
	if !f.Deleted {
		filter["deleted_at"] = nil
	}

	// id
	if f.Id != nil && !f.Id.IsZero() {
		filter["_id"] = f.Id
	}

	// refresh token
	if f.RefreshToken != "" {
		filter["refresh_token"] = f.RefreshToken
	}

	return filter
}
