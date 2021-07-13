package bd

import (
	"context"
	"time"

	model "github.com/jacobd39/admsys/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//InsertNewProfile inserta un nuevo perfil en la base de datos
func InsertNewProfile(p model.Profile) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("admsys")
	col := db.Collection("profiles")

	result, err := col.InsertOne(ctx, p)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjID.String(), true, nil
}
