package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type RefreshTokenModel struct {
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
