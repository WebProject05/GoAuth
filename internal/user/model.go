package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id, omitempty"`
	Email        string             `bson:"email" json:"email"`
	PasswordHash string             `bson:"passwordhash" json:"-"`
	Role         string             `bson:"role" json:"role"`
	CreatedAt    time.Time          `bson:"createdat" json:"createdat"`
	UpdatedAt    time.Time          `bson:"updatedat" json:"updatedat"`
}

type PublicUser struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}



func ToPublic(u User) PublicUser {
	return PublicUser{
		ID: u.ID.Hex(),
		Email: u.Email,
		Role: u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}