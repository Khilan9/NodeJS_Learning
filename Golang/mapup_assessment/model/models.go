package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Location struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty"`
	Address   string             `json:"address,omitempty"`
	Latitude  float64            `json:"latitude,omitempty"`
	Longitude float64            `json:"longitude,omitempty"`
	Category  string             `json:"category,omitempty"`
}
