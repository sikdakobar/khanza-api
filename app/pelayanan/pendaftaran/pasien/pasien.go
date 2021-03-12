package pasien

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

// Pasien is
type Pasien struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	NIK           int                `bson:"nik,omitempty"`
	Nama          string             `bson:"nama,omitempty"`
	DOB           string             `bson:"dob,omitempty"`
	POB           string             `bson:"pob,omitempty"`
	Age           int                `bson:"age,omitempty"`
	Jenis_Kelamin string             `bson:"jenis_kelamin,omitempty"`
	GolDarah      string             `bson:"gol_darah,omitempty"`
	Alamat        []Alamat           `bson:"alamat,omitempty"`
	Rekam_Medis   []RekamMedis       `bson:"rekam_medis,omitempty"`
	CreatedAt     primitive.DateTime `bson:"createdat,omitempty"`
}

type Alamat struct {
	Jalan     string             `bson:"jalan,omitempty"`
	No        uint8              `bson:"no,omitempty"`
	RT        uint8              `bson:"rt,omitempty"`
	RW        uint8              `bson:"rw,omitempty"`
	Kelurahan string             `bson:"kelurahan,omitempty"`
	Kecamatan string             `bson:"kecamatan,omitempty"`
	Kabupaten string             `bson:"kabupaten,omitempty"`
	Provinsi  string             `bson:"provinsi,omitempty"`
	CreatedAt primitive.DateTime `bson:"createdat,omitempty"`
}

type Biometrik struct {
}

type RekamMedis struct {
}

// Index is
func Index(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	db, err := db.MongoDB()

	var pasien []Pasien

	query, err := db.Collection("pasien").Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err.Error())
	}

	defer query.Close(context.Background())

	for query.Next(context.Background()) {
		var data Pasien
		query.Decode(&data)
		pasien = append(pasien, data)
	}

	json.NewEncoder(res).Encode(pasien)
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
	var pasien Pasien
	db.Collection("pasien").FindOne(context.Background(), bson.M{"_id": id}).Decode(&pasien)
	json.NewEncoder(res).Encode(pasien)
}

// Store is
func Store(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	db, err := db.MongoDB()

	var pasien Pasien
	json.NewDecoder(req.Body).Decode(&pasien)
	// data := bson.D{
	// 	{Key: "no_rm", Value: pasien.NoRM},
	// 	{Key: "name", Value: pasien.Name},
	// 	{Key: "dob", Value: pasien.DOB},
	// 	{Key: "pob", Value: pasien.POB},
	// 	{Key: "age", Value: pasien.Age},
	// 	{Key: "jenis_kelamin", Value: pasien.JenisKelamin},
	// 	{Key: "gol_darah", Value: pasien.GolDarah},
	// 	{Key: "alamat", Value: pasien.Alamat},
	// 	{Key: "createdat", Value: time.Now()}}

	db.Collection("pasien").InsertOne(context.Background(), pasien)
	if err != nil {
		log.Fatal(err.Error())
	}

	json.NewEncoder(res).Encode(pasien)
}

// Update is
func Update(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	db, err := db.MongoDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	var pasien Pasien
	json.NewDecoder(req.Body).Decode(&pasien)
	params := mux.Vars(req)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	data := bson.M{"$set": bson.M{
		"name":          pasien.Nama,
		"nik":           pasien.NIK,
		"dob":           pasien.DOB,
		"pob":           pasien.POB,
		"age":           pasien.Age,
		"jenis_kelamin": pasien.Jenis_Kelamin,
		"gol_darah":     pasien.GolDarah,
	}}

	db.Collection("pasien").FindOneAndUpdate(context.Background(), Pasien{ID: id}, data).Decode(&pasien)
	json.NewEncoder(res).Encode(pasien)
}

// Destroy is
func Destroy(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	db, err := db.MongoDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	var pasien Pasien

	params := mux.Vars(req)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	db.Collection("pasien").FindOneAndDelete(context.Background(), Pasien{ID: id})
	json.NewEncoder(res).Encode(pasien)
}

func AlamatStore(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	db, err := db.MongoDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	var alamat Alamat
	json.NewDecoder(req.Body).Decode(&alamat)
	params := mux.Vars(req)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	data := bson.M{"$push": bson.M{"alamat": bson.M{"Jalan": alamat.Jalan, "No": alamat.No, "RT": alamat.RT, "RW": alamat.RW, "Kelurahan": alamat.Kelurahan, "Kecamatan": alamat.Kecamatan, "Kabupaten": alamat.Kabupaten, "Provinsi": alamat.Provinsi}}}
	db.Collection("pasien").FindOneAndUpdate(context.Background(), Pasien{ID: id}, data).Decode(&alamat)
	json.NewEncoder(res).Encode(alamat)
}

func AlamatUpdate(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	db, err := db.MongoDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	var alamat Alamat
	json.NewDecoder(req.Body).Decode(&alamat)
	params := mux.Vars(req)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	index := params["index"]
	data := bson.M{"$set": bson.M{"alamat" + "." + index: bson.M{
		"Jalan":     alamat.Jalan,
		"No":        alamat.No,
		"RT":        alamat.RT,
		"RW":        alamat.RW,
		"Kelurahan": alamat.Kelurahan,
		"Kecamatan": alamat.Kecamatan,
		"Kabupaten": alamat.Kabupaten,
		"Provinsi":  alamat.Provinsi,
	}},
	}

	db.Collection("pasien").FindOneAndUpdate(context.Background(), Pasien{ID: id}, data).Decode(&alamat)
	json.NewEncoder(res).Encode(alamat)
}

func RekamMedisIndex(res http.ResponseWriter, req *http.Request) {
	// res.Header().Set("Content-type", "application/json")
	// db, err := db.MongoDB()
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// var rekam_medis RekamMedis
	// json.NewDecoder(req.Body).Decode(&rekam_medis)
	// params := mux.Vars(req)
	// id, _ := primitive.ObjectIDFromHex(params["id"])

	// db.Collection("pasien").FindOne(context.Background(), bson.M{"_id": id}).Decode(&rekam_medis)
	// json.NewEncoder(res).Encode(rekam_medis)
}

func RekamMedisStore(res http.ResponseWriter, req *http.Request) {

}

func RekamMedisUpdate(res http.ResponseWriter, req *http.Request) {

}
