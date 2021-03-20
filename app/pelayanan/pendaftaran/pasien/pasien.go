package pasien

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

// Pasien is
type Pasien struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	NIK           int                `bson:"nik,omitempty"`
	No_RM         int                `bson:"no_rm,omitempty"`
	Nama          string             `bson:"nama,omitempty"`
	DOB           primitive.DateTime `bson:"dob,omitempty"`
	POB           string             `bson:"pob,omitempty"`
	Jenis_Kelamin string             `bson:"jenis_kelamin,omitempty"`
	GolDarah      string             `bson:"gol_darah,omitempty"`
	Email         string             `bson:"email,omitempty"`
	Password      string             `bson:"password,omitempty"`
	Token         string             `bson:"token,omitempty"`
	Alamat        []Alamat           `bson:"alamat,omitempty"`
	Rekam_Medis   []Rekam_Medis      `bson:"rekam_medis,omitempty"`
	Age           int
	CreatedAt     primitive.DateTime `bson:"createdat,omitempty"`
}

type Alamat struct {
	Jalan          string             `bson:"jalan,omitempty"`
	No             uint8              `bson:"no,omitempty"`
	RT             uint8              `bson:"rt,omitempty"`
	RW             uint8              `bson:"rw,omitempty"`
	Desa_Kelurahan string             `bson:"desa_kelurahan,omitempty"`
	Kecamatan      string             `bson:"kecamatan,omitempty"`
	Kabupaten      string             `bson:"kabupaten,omitempty"`
	Provinsi       string             `bson:"provinsi,omitempty"`
	CreatedAt      primitive.DateTime `bson:"createdat,omitempty"`
}

type Biometrik struct {
	Face        string
	Fingerprint string
	Iris        string
	CreatedAt   primitive.DateTime `bson:"createdat,omitempty"`
}

type Rekam_Medis struct {
	ICD_Code          string             `bson:"icd_code,omitempty"`
	Poli              string             `bson:"poli,omitempty"`
	Diagnosa_Penyakit string             `bson:"diagnosa_penyakit,omitempty"`
	Keluhan           string             `bson:"keluhan,omitempty"`
	Pemeriksaan_Fisik string             `bson:"pemeriksaan_fisik,omitempty"`
	Pemeriksaan_Lab   string             `bson:"pemeriksaan_lab,omitempty"`
	Perawatan         string             `bson:"perawatan,omitempty"`
	CreatedAt         primitive.DateTime `bson:"createdat,omitempty"`
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

	for i := range pasien {
		pasien[i].Age = time.Now().Year() - pasien[i].Age
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

	pasien.Age = time.Now().Year() - pasien.Age

	json.NewEncoder(res).Encode(pasien)
}

// Store is
func Store(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	db, err := db.MongoDB()

	var pasien Pasien
	json.NewDecoder(req.Body).Decode(&pasien)

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
	data := bson.M{"$push": bson.M{"alamat": bson.M{"Jalan": alamat.Jalan, "No": alamat.No, "RT": alamat.RT, "RW": alamat.RW, "Desa_Kelurahan": alamat.Desa_Kelurahan, "Kecamatan": alamat.Kecamatan, "Kabupaten": alamat.Kabupaten, "Provinsi": alamat.Provinsi, "CreatedAt": alamat.CreatedAt}}}
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
		"Jalan":          alamat.Jalan,
		"No":             alamat.No,
		"RT":             alamat.RT,
		"RW":             alamat.RW,
		"Desa_Kelurahan": alamat.Desa_Kelurahan,
		"Kecamatan":      alamat.Kecamatan,
		"Kabupaten":      alamat.Kabupaten,
		"Provinsi":       alamat.Provinsi,
	}},
	}

	db.Collection("pasien").FindOneAndUpdate(context.Background(), Pasien{ID: id}, data).Decode(&alamat)
	json.NewEncoder(res).Encode(alamat)
}

func RekamMedisStore(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	db, err := db.MongoDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	var rekam_medis Rekam_Medis
	json.NewDecoder(req.Body).Decode(&rekam_medis)
	params := mux.Vars(req)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	data := bson.M{"$push": bson.M{"rekam_medis": bson.M{
		"ICD_Code":          rekam_medis.ICD_Code,
		"Poli":              rekam_medis.Poli,
		"Diagnosa_Penyakit": rekam_medis.Diagnosa_Penyakit,
		"Keluhan":           rekam_medis.Keluhan,
		"Pemeriksaan_Fisik": rekam_medis.Pemeriksaan_Fisik,
		"Pemeriksaan_Lab":   rekam_medis.Pemeriksaan_Lab,
		"Perawatan":         rekam_medis.Perawatan,
		"CreatedAt":         rekam_medis.CreatedAt}}}
	db.Collection("pasien").FindOneAndUpdate(context.Background(), Pasien{ID: id}, data).Decode(&rekam_medis)
	json.NewEncoder(res).Encode(rekam_medis)
}

func RekamMedisUpdate(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	db, err := db.MongoDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	var rekam_medis Rekam_Medis
	json.NewDecoder(req.Body).Decode(&rekam_medis)
	params := mux.Vars(req)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	index := params["index"]
	data := bson.M{"$set": bson.M{"rekam_medis" + "." + index: bson.M{
		"ICD_Code":          rekam_medis.ICD_Code,
		"Poli":              rekam_medis.Poli,
		"Diagnosa_Penyakit": rekam_medis.Diagnosa_Penyakit,
		"Keluhan":           rekam_medis.Keluhan,
		"Pemeriksaan_Fisik": rekam_medis.Pemeriksaan_Fisik,
		"Pemeriksaan_Lab":   rekam_medis.Pemeriksaan_Lab,
		"Perawatan":         rekam_medis.Perawatan,
	}},
	}

	db.Collection("pasien").FindOneAndUpdate(context.Background(), Pasien{ID: id}, data).Decode(&rekam_medis)
	json.NewEncoder(res).Encode(rekam_medis)
}
