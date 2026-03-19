package bd

import (
	"context"
	"fmt"

	"github.com/harolpg17/twitterGo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func LeoTweets(ID string, pagina int64) ([]*models.DevuelvoTweet, bool) {
	ctx := context.TODO()

	db := MongoCN.Database(DatabaseName)
	col := db.Collection("tweet")

	var resultado []*models.DevuelvoTweet

	condicion := bson.M{
		"userid": ID,
	}

	opciones := options.Find()
	opciones.SetLimit(20)
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}})
	opciones.SetSkip((pagina - 1) * 20)

	cursor, err := col.Find(ctx, condicion, opciones)
	if err != nil {
		fmt.Println("Error al leer los tweets:", err)
		return resultado, false
	}

	for cursor.Next(ctx) {
		var registro models.DevuelvoTweet
		err := cursor.Decode(&registro)
		if err != nil {
			fmt.Println("Error al decodificar el tweet:", err)
			return resultado, false
		}

		resultado = append(resultado, &registro)
	}

	return resultado, true
}
