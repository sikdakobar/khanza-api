package pegawai

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"simpus/db"

	// "time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Pegawai is
type Pegawai struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Nama          string             `bson:"nama,omitempty"`
	Jabatan       string             `bson:"jabatan,omitempty"`
	Pendidikan    string             `bson:"pendidikan,omitempty"`
	Status        string             `bson:"status,omitempty"`
	Jenis_Kelamin string             `bson:"jenis_kelamin,omitempty"`
	POB           string             `bson:"pob,omitempty"`
	DOB           string             `bson:"dob,omitempty"`
	TMT           []TMT              `bson:"tmt,omitempty"`
}

type TMT struct {
	NoSK      string             `bson:"no_sk,omitempty"`
	CreatedAt primitive.DateTime `bson:"createdat,omitempty"`
}

// Index is
func Index(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	db, err := db.MongoDB()

	var pegawai []Pegawai

	query, err := db.Collection("pegawai").Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err.Error())
	}

	defer query.Close(context.Background())

	for query.Next(context.Background()) {
		var data Pegawai
		query.Decode(&data)
		pegawai = append(pegawai, data)
	}

	json.NewEncoder(res).Encode(pegawai)
}

// Show is
func Show(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	db, err := db.MongoDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	params := mux.Vars(req)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var pegawai Pegawai
	db.Collection("pegawai").FindOne(context.Background(), bson.M{"_id": id}).Decode(&pegawai)
	json.NewEncoder(res).Encode(pegawai)
}

// Store is
func Store(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	db, err := db.MongoDB()

	var pegawai Pegawai
	json.NewDecoder(req.Body).Decode(&pegawai)

	db.Collection("pegawai").InsertOne(context.Background(), pegawai)
	if err != nil {
		log.Fatal(err.Error())
	}

	json.NewEncoder(res).Encode(pegawai)
}

// Update is
func Update(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	db, err := db.MongoDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	var pegawai Pegawai
	json.NewDecoder(req.Body).Decode(&pegawai)
	params := mux.Vars(req)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	data := bson.M{"$set": bson.M{
		"nama":          pegawai.Nama,
		"jabatan":       pegawai.Jabatan,
		"pendidikan":    pegawai.Pendidikan,
		"status":        pegawai.Status,
		"jenis_kelamin": pegawai.Jenis_Kelamin,
		"pob":           pegawai.POB,
		"dob":           pegawai.DOB,
		"tmt":           pegawai.TMT,
	}}

	db.Collection("pegawai").FindOneAndUpdate(context.Background(), Pegawai{ID: id}, data).Decode(&pegawai)
	json.NewEncoder(res).Encode(pegawai)
}

// Destroy is
func Destroy(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	db, err := db.MongoDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	var pegawai Pegawai

	params := mux.Vars(req)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	db.Collection("pegawai").FindOneAndDelete(context.Background(), Pegawai{ID: id})
	json.NewEncoder(res).Encode(pegawai)

}

func TMTUpdate(res http.ResponseWriter, req *http.Request) {

}
