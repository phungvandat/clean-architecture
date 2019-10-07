package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User model
type User struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Username string             `json:"username,omitempty"`
	Fullname string             `json:"fullname,omitempty"`
	Password string             `json:"-"`
	Email    string             `json:"email,omitempty"`
	Phone    string             `json:"phone,omitempty"`
	Role     string             `json:"role,omitempty"`
}
