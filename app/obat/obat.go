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
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Nama     string             `bson:"nama,omitempty"`
	Kategori string             `bson:"kategori,omitempty"`
	Golongan string             `bson:"golongan,omitempty"`
	Batch    []Batch            `bson:"batch_obat,omitempty"`
}

type Batch struct {
	Merek      string `bson:"merek,omitempty"`
	Dosis      string `bson:"dosis,omitempty"`
	Harga      string `bson:"harga,omitempty"`
	Bentuk     string `bson:"bentuk,omitempty"`
	Supplier   string `bson:"supplier,omitempty"`
	Expired    string `bson:"expired,omitempty"`
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
		"nama":     obat.Nama,
		"kategori": obat.Kategori,
		"golongan": obat.Golongan,
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

func BatchObatStore(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	db, err := db.MongoDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	var batch_obat Batch
	json.NewDecoder(req.Body).Decode(&batch_obat)
	params := mux.Vars(req)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	data := bson.M{"$push": bson.M{"batch_obat": bson.M{
		"Merek":      batch_obat.Merek,
		"Dosis":      batch_obat.Dosis,
		"Harga":      batch_obat.Harga,
		"Bentuk":     batch_obat.Bentuk,
		"Supplier":   batch_obat.Supplier,
		"Expired":    batch_obat.Expired,
		"Date_Entry": batch_obat.Date_Entry}}}
	db.Collection("obat").FindOneAndUpdate(context.Background(), Obat{ID: id}, data).Decode(&batch_obat)
	json.NewEncoder(res).Encode(batch_obat)
}

func BatchObatUpdate(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	db, err := db.MongoDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	var batch_obat Batch
	json.NewDecoder(req.Body).Decode(&batch_obat)
	params := mux.Vars(req)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	index := params["index"]
	data := bson.M{"$push": bson.M{"batch_obat" + "." + index: bson.M{
		"Merek":      batch_obat.Merek,
		"Dosis":      batch_obat.Dosis,
		"Harga":      batch_obat.Harga,
		"Bentuk":     batch_obat.Bentuk,
		"Supplier":   batch_obat.Supplier,
		"Expired":    batch_obat.Expired,
		"Date_Entry": batch_obat.Date_Entry}}}
	db.Collection("obat").FindOneAndUpdate(context.Background(), Obat{ID: id}, data).Decode(&batch_obat)
	json.NewEncoder(res).Encode(batch_obat)
}

func BatchObatDestroy(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	db, err := db.MongoDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	var batch_obat Batch
	json.NewDecoder(req.Body).Decode(&batch_obat)
	params := mux.Vars(req)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	index := params["index"]
	data := bson.M{"$pop": bson.M{"batch_obat" + "." + index: bson.M{}}}
	db.Collection("obat").FindOneAndUpdate(context.Background(), Obat{ID: id}, data).Decode(&batch_obat)
	json.NewEncoder(res).Encode(batch_obat)
}
