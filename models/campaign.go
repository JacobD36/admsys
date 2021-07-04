package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Campaign ---> Define la estructura de la tabla Campañas
type Campaign struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Desc        string             `bson:"desc" json:"desc"`
	Status      int                `bson:"status" json:"status"`
	Responsable string             `bson:"responsable" json:"responsable"`
	Phone       string             `bson:"phone" json:"phone"`
	CreateAt    time.Time          `bson:"createAt" json:"createAt,omitempty"`
	ModifiedAt  time.Time          `bson:"modifiedAt" json:"modifiedAt,omitempty"`
}
