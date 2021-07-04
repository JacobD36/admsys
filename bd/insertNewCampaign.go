package bd

import (
	"context"
	"time"

	model "github.com/jacobd39/admsys/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//InserNewCampaign inserta una nueva campaña en la base de datos
func InserNewCampaign(c model.Campaign) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("admsys")
	col := db.Collection("campaigns")

	result, err := col.InsertOne(ctx, c)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjID.String(), true, nil
}
