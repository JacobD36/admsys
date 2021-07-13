package bd

import (
	"context"
	"time"

	model "github.com/jacobd39/admsys/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//InsertNewPosition inserta un nuevo puesto en la base de datos
func InsertNewPosition(p model.Positions) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("admsys")
	col := db.Collection("positions")

	p.ID_position = primitive.ObjectID(p.ID_position)
	p.ID_profile = primitive.ObjectID(p.ID_profile)

	result, err := col.InsertOne(ctx, p)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjID.String(), true, nil
}
