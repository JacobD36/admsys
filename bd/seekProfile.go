package bd

import (
	"context"
	"fmt"
	"time"

	model "github.com/jacobd39/admsys/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SeekProfile busca un registro de perfil específico en la base de datos
func SeekProfile(ID string) (model.Profile, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("admsys")
	col := db.Collection("profiles")

	var profileData model.Profile

	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condition).Decode(&profileData)

	if err != nil {
		fmt.Println("Registro no encontrado " + err.Error())
		return profileData, err
	}

	return profileData, nil
}
