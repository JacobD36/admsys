package bd

import (
	"context"
	"fmt"
	"time"

	model "github.com/jacobd39/admsys/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//GetPositions devuelve el listado de los perfiles activos
func GetPositions(page int64, search string) ([]*model.Positions, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("admsys")
	col := db.Collection("positions")

	var results []*model.Positions

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 10)
	findOptions.SetSort(bson.M{"createAt": -1})
	findOptions.SetLimit(10)

	query := bson.M{
		"title": bson.M{"$regex": `(?i)` + search},
	}

	cur, err := col.Find(ctx, query, findOptions)

	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	for cur.Next(ctx) {
		var s model.Positions

		err := cur.Decode(&s)

		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}

		results = append(results, &s)
	}

	err = cur.Err()

	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	cur.Close(ctx)
	return results, true
}