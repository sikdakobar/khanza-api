package harga

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"simpus/db"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Harga is
type Harga struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name,omitempty"`
	Cost      float32            `bson:"cost,omitempty"`
	CreatedAt primitive.DateTime `bson:"createdat,omitempty"`
}

// Index is
func Index(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	db, err := db.MongoDB()

	var harga []Harga

	query, err := db.Collection("harga").Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err.Error())
	}

	defer query.Close(context.Background())

	for query.Next(context.Background()) {
		var data Harga
		query.Decode(&data)
		harga = append(harga, data)
	}

	json.NewEncoder(res).Encode(harga)
}

// Store is
func Store(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	db, err := db.MongoDB()

	var harga Harga

	json.NewDecoder(req.Body).Decode(&harga)
	data := bson.D{

		{Key: "name", Value: harga.Name},
		{Key: "cost", Value: harga.Cost},
		{Key: "createdat", Value: time.Now()},
	}

	db.Collection("harga").InsertOne(context.Background(), data)

	if err != nil {
		log.Fatal(err.Error())
	}

	json.NewEncoder(res).Encode(harga)

}

// Update is
func Update(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	db, err := db.MongoDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	var harga Harga
	json.NewDecoder(req.Body).Decode(&harga)
	params := mux.Vars(req)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	data := bson.D{
		{"$set", bson.D{
			{Key: "name", Value: harga.Name},
			{Key: "cost", Value: harga.Cost},
		}}}

	db.Collection("harga").FindOneAndUpdate(context.Background(), Harga{ID: id}, data).Decode(&harga)
	json.NewEncoder(res).Encode(harga)
}

// Destroy is
func Destroy(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	db, err := db.MongoDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	var harga Harga

	params := mux.Vars(req)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	db.Collection("harga").FindOneAndDelete(context.Background(), Harga{ID: id})
	json.NewEncoder(res).Encode(harga)
}
