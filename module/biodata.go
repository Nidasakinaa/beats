package beats

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

// insertonedoc

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertBiodata(nama string, jk string, agama string, prodi string, datadiri interface{}, alamat interface{}) interface{} {
	biodata := map[string]interface{}{
		"bio_id":   primitive.NewObjectID(),
		"nama":     nama,
		"jk":       jk,
		"agama":    agama,
		"prodi":    prodi,
		"datadiri": datadiri,
		"alamat":   alamat,
	}
	return InsertOneDoc("ATS", "biodata", biodata)
}

func GetAllBiodata() ([]interface{}, error) {
	collection := MongoConnect("ATS").Collection("biodata")
	filter := bson.M{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllBiodata:", err)
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var biodata []interface{}
	if err := cursor.All(context.TODO(), &biodata); err != nil {
		fmt.Println("GetAllBiodata:", err)
		return nil, err
	}
	return biodata, nil
}
