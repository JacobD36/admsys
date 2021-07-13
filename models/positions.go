package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Positions ---> Define la estructura de la tabla posiciones (Puestos)
type Positions struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	ID_position primitive.ObjectID `bson:"id_position,omitempty" json:"id_position,omitempty"`
	ID_profile  primitive.ObjectID `bson:"id_profile,omitempty" json:"id_profile,omitempty"`
	CreateAt    time.Time          `bson:"createAt" json:"createAt,omitempty"`
	ModifiedAt  time.Time          `baon:"modifiedAt" json:"modifiedAt,omitempty"`
}
