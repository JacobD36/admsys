package bd

import (
	"context"
	"time"

	model "github.com/jacobd39/admsys/models"
	"go.mongodb.org/mongo-driver/bson"
)

//UserVerification recibe un codUser de parámetro y verifica si ya existe en la BD
func UserVerification(dni string) (model.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("admsys")
	col := db.Collection("users")

	condition := bson.M{
		"dni": dni,
	}

	var result model.User

	err := col.FindOne(ctx, condition).Decode(&result)

	ID := result.ID.Hex()

	if err != nil {
		return result, false, ID
	}

	return result, true, ID
}
