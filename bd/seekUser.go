package bd

import (
	"context"
	"fmt"
	"time"

	model "github.com/jacobd39/admsys/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SeekUser busca un registro en la base de datos
func SeekUser(ID string) (model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("admsys")
	col := db.Collection("users")

	var usrData model.User

	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condition).Decode(&usrData)

	usrData.Password = ""

	if err != nil {
		fmt.Println("Registro no encontrado " + err.Error())
		return usrData, err
	}

	return usrData, nil
}
