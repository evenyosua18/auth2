package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserModel struct {
	Id       primitive.ObjectID `bson:"_id"`
	Email    string             `bson:"email"`
	Phone    string             `bson:"phone"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
	IsActive bool               `bson:"is_active"`

	CreatedAt time.Time  `bson:"created_at"`
	UpdatedAt time.Time  `bson:"updated_at"`
	DeletedAt *time.Time `bson:"deleted_at"`
}
