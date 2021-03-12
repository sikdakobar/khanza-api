package obat

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"simpus/db"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Obat struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Nama  string             `bson:"nama,omitempty"`
	Batch []string           `bson:"batch,omitempty"`
}

type Batch struct {
	Merek      string `bson:"merek,omitempty"`
	Jenis      string `bson:"jenis,omitempty"`
	Golongan   string `bson:"golongan,omitempty"`
	Harga      string `bson:"harga,omitempty"`
	Supplier   string `bson:"supplier,omitempty"`
	Date_Entry string `bson:"date_entry,omitempty"`
}

func Index(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	db, err := db.MongoDB()

	var obat []Obat

	query, err := db.Collection("obat").Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err.Error())
	}

	defer query.Close(context.Background())

	for query.Next(context.Background()) {
		var data Obat
		query.Decode(&data)
		obat = append(obat, data)
	}

	json.NewEncoder(res).Encode(obat)
}

func Store(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	db, err := db.MongoDB()

	var obat Obat
	json.NewDecoder(req.Body).Decode(&obat)

	db.Collection("obat").InsertOne(context.Background(), obat)
	if err != nil {
		log.Fatal(err.Error())
	}

	json.NewEncoder(res).Encode(obat)
}

func Update(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	db, err := db.MongoDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	var obat Obat
	json.NewDecoder(req.Body).Decode(&obat)
	params := mux.Vars(req)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	data := bson.M{"$set": bson.M{
		"name": obat.Nama,
	}}

	db.Collection("obat").FindOneAndUpdate(context.Background(), Obat{ID: id}, data).Decode(&obat)
	json.NewEncoder(res).Encode(obat)
}

func Destroy(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	db, err := db.MongoDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	var obat Obat

	params := mux.Vars(req)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	db.Collection("obat").FindOneAndDelete(context.Background(), Obat{ID: id})
	json.NewEncoder(res).Encode(obat)
}
