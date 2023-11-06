package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type AccessTokenModel struct {
	Id        primitive.ObjectID `bson:"_id"`
	ExpiredAt time.Time          `bson:"expired_at"`
	UserId    primitive.ObjectID `bson:"user_id"`

	CreatedAt time.Time  `bson:"created_at"`
	UpdatedAt time.Time  `bson:"updated_at"`
	DeletedAt *time.Time `bson:"deleted_at"`
}
