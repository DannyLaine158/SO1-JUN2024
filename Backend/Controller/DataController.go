package Controller

import (
	"Backend/Instance"
	"Backend/Model"
	"context"
	"log"
)

func InsertData(nameCol string, dataParam string) {
	collection := Instance.Mg.Db.Collection(nameCol)
	doc := Model.Data{Percent: dataParam}

	_, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Fatal(err)
	}
}
