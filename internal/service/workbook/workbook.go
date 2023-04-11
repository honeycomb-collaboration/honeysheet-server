package workbook

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"honeysheet-server/internal/model"
)

type Workbook struct {
	*model.Workbook
}

const WorkbookCollection = "Workbook"

func (workbook *Workbook) GetByID(ID string) error {
	findOptions := options.FindOne()
	findOptions.SetMaxTime(500)
	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}

	err = model.MongoDBClient.Collection(WorkbookCollection).
		FindOne(context.TODO(), bson.M{"_id": objID, "IsDeleted": bson.M{"$ne": true}}, findOptions).
		Decode(workbook)

	return err

}
