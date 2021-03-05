package antrian

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"simpus/db"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Antrian struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	NIK       int                `bson:"nik,omitempty"`
	Poli      string             `bson:"poli,omitempty"`
	Date      string             `bson:"date,omitempty"`
	CreatedAt primitive.DateTime `bson:"createdat,omitempty"`
}

func Index(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	db, err := db.MongoDB()

	var antrian []Antrian

	query, err := db.Collection("antrian").Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err.Error())
	}

	defer query.Close(context.Background())

	for query.Next(context.Background()) {
		var data Antrian
		query.Decode(&data)
		antrian = append(antrian, data)
	}

	json.NewEncoder(res).Encode(antrian)
}

func Store(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	db, err := db.MongoDB()

	var antrian Antrian
	json.NewDecoder(req.Body).Decode(&antrian)

	db.Collection("antrian").InsertOne(context.Background(), antrian)
	if err != nil {
		log.Fatal(err.Error())
	}

	json.NewEncoder(res).Encode(antrian)
}

func ListAntrian(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	db, err := db.MongoDB()

	var antrian []Antrian

	query, err := db.Collection("antrian").Find(context.Background(), bson.M{"date": time.Now().Format("1900-01-01")})
	if err != nil {
		log.Fatal(err.Error())
	}

	defer query.Close(context.Background())

	for query.Next(context.Background()) {
		var data Antrian
		query.Decode(&data)
		antrian = append(antrian, data)
	}

	json.NewEncoder(res).Encode(antrian)
}
