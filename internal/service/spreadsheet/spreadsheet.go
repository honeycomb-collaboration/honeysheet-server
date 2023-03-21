package spreadsheet

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"honeysheet-server/internal/model"
)

type SpreadSheet struct {
	*model.SpreadSheet
}

var SpreadSheetCollection = "SpreadSheet"

func (ss *SpreadSheet) GetByID(ID string) error {
	findOptions := options.FindOne()
	findOptions.SetMaxTime(500)
	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}

	err = model.MongoDBClient.Collection(SpreadSheetCollection).
		FindOne(context.TODO(), bson.M{"_id": objID, "IsDeleted": bson.M{"$ne": true}}, findOptions).
		Decode(ss)

	return err

}
