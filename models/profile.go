package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Profile ---> Define la estructura de la tabla Perfiles
type Profile struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title      string             `bson:"title" json:"title"`
	CreateAt   time.Time          `bson:"createAt" json:"createAt,omitempty"`
	ModifiedAt time.Time          `baon:"modifiedAt" json:"modifiedAt,omitempty"`
}
