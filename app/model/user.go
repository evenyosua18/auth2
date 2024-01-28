package model

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	UserModel struct {
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

	UserFilter struct {
		Id         *primitive.ObjectID
		Email      string
		Username   string
		Phone      string
		IsActive   bool
		CheckExist bool
		Value      string // for login, email or username or phone

		Deleted bool
	}
)

func (m UserModel) SetCreated() UserModel {
	if m.Id.IsZero() {
		m.Id = primitive.NewObjectID()
	}

	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()

	return m
}

func (f UserFilter) Filter() bson.M {
	filter := bson.M{}
	or := bson.A{}

	// deleted
	if !f.Deleted {
		filter["deleted_at"] = nil
	}

	if f.CheckExist {
		// phone
		if f.Phone != "" {
			or = append(or, bson.M{"phone": f.Phone})
		}

		// email
		if f.Email != "" {
			or = append(or, bson.M{"email": f.Email})
		}

		// username
		if f.Username != "" {
			or = append(or, bson.M{"username": f.Username})
		}
	} else {
		// phone
		if f.Phone != "" {
			filter["phone"] = primitive.Regex{Pattern: f.Phone, Options: "i"}
		}

		// email
		if f.Email != "" {
			filter["email"] = primitive.Regex{Pattern: f.Email, Options: "i"}
		}

		// username
		if f.Username != "" {
			filter["username"] = f.Username
		}
	}

	// phone or email or username, used by login
	if f.Value != "" {
		or = append(or, bson.M{"username": f.Value}, bson.M{"email": f.Value}, bson.M{"phone": f.Value})
	}

	// id
	if f.Id != nil && !f.Id.IsZero() {
		filter["_id"] = f.Id
	}

	// is active
	if f.IsActive {
		filter["is_active"] = f.IsActive
	}

	// add or if exist
	if len(or) != 0 {
		filter["$or"] = or
	}

	return filter
}
