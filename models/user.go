package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User ---> Degine la estructura de la tabla usuarios
type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Dni   	   string             `bson:"dni" json:"dni"`
	Email	   string			  `bson:"email" json:"email"`
	IDProfile  int                `bson:"idProfile" json:"idProfile"`
	Status     int                `bson:"status" json:"status"`
	Name1      string             `bson:"name1" json:"name1"`
	Name2      string             `bson:"name2" json:"name2,omitempty"`
	LastName1  string             `bson:"lastName1" json:"lastName1"`
	LastName2  string             `bson:"lastName2" json:"lastName2,omitempty"`
	Password   string             `bson:"password" json:"password"`
	BirthDate  time.Time		  `bson:"birthDate" json:"birthDate"`
	CreateAt   time.Time          `bson:"createAt" json:"createAt,omitempty"`
	ModifiedAt time.Time          `bson:"modifiedAt" json:"modifiedAt,omitempty"`
}
