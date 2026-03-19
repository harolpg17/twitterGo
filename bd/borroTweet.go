package bd

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BorroTweet(IDTweet string, IDUsuario string) error {
	ctx := context.TODO()

	db := MongoCN.Database(DatabaseName)
	col := db.Collection("tweet")

	ObjID, err := primitive.ObjectIDFromHex(IDTweet)
	if err != nil {
		return fmt.Errorf("invalid tweet ID")
	}

	condicion := bson.M{
		"_id":    ObjID,
		"userid": IDUsuario,
	}

	_, err = col.DeleteOne(ctx, condicion)
	if err != nil {
		return fmt.Errorf("error deleting tweet")
	}

	return nil
}
