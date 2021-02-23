package keuangan

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"simpus/db"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ChartAccount is
type ChartAccount struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Name       string             `bson:"name,omitempty"`
	Type       string             `bson:"type,omitempty"`
	CreatedAt  primitive.DateTime `bson:"createdat,omitempty"`
	SubAccount []SubAccount       `bson:"subaccount,omitempty"`
}

// SubAccount is
type SubAccount struct {
	Kode    string `bson:"kode,omitempty"`
	Name    string `bson:"name,omitempty"`
	Balance string `bson:"balance,omitempty"`
}

// Index is
func IndexCOA(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	db, err := db.MongoDB()

	var account []ChartAccount

	query, err := db.Collection("coa").Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err.Error())
	}

	defer query.Close(context.Background())

	for query.Next(context.Background()) {
		var data ChartAccount
		query.Decode(&data)
		account = append(account, data)
	}

	json.NewEncoder(res).Encode(account)
}

// Show is
func ShowCOA(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	db, err := db.MongoDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	params := mux.Vars(req)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var account ChartAccount
	db.Collection("coa").FindOne(context.Background(), bson.M{"_id": id}).Decode(&account)
	json.NewEncoder(res).Encode(account)
}

// Store is
func StoreCOA(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	db, err := db.MongoDB()

	var account ChartAccount
	json.NewDecoder(req.Body).Decode(&account)
	data := bson.D{
		{Key: "name", Value: account.Name},
		{Key: "type", Value: account.Type},
		{Key: "createdat", Value: time.Now()},
		{Key: "subaccount", Value: account.SubAccount}}

	db.Collection("coa").InsertOne(context.Background(), data)
	if err != nil {
		log.Fatal(err.Error())
	}

	json.NewEncoder(res).Encode(account)
}

// Update is
func UpdateCOA(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	db, err := db.MongoDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	var account ChartAccount
	json.NewDecoder(req.Body).Decode(&account)
	params := mux.Vars(req)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	data := bson.D{
		{"$set", bson.D{
			{Key: "name", Value: account.Name},
			{Key: "type", Value: account.Type},
		}}}

	db.Collection("coa").FindOneAndUpdate(context.Background(), ChartAccount{ID: id}, data).Decode(&account)
	json.NewEncoder(res).Encode(account)
}

// Destroy is
func DestroyCOA(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")
	db, err := db.MongoDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	var account ChartAccount

	params := mux.Vars(req)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	db.Collection("coa").FindOneAndDelete(context.Background(), ChartAccount{ID: id})
	json.NewEncoder(res).Encode(account)

}

// StoreSubAccount is
func StoreSubAccountCOA(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	db, err := db.MongoDB()

	var account ChartAccount
	json.NewDecoder(req.Body).Decode(&account)
	params := mux.Vars(req)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	data := bson.M{"$push": bson.M{"subaccount": bson.M{"$each": account.SubAccount}}}

	db.Collection("coa").FindOneAndUpdate(context.Background(), ChartAccount{ID: id}, data).Decode(&account.SubAccount)
	if err != nil {
		log.Fatal(err.Error())
	}

	json.NewEncoder(res).Encode(account.SubAccount)
}

// DestroySubAccount is
func DestroySubAccountCOA(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	db, err := db.MongoDB()

	var account ChartAccount
	json.NewDecoder(req.Body).Decode(&account)
	params := mux.Vars(req)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	index, _ := strconv.Atoi(params["index"])
	data := bson.M{
		"$pull": bson.M{
			"subaccount": account.SubAccount[index],
		}}

	db.Collection("coa").FindOneAndUpdate(context.Background(), ChartAccount{ID: id}, data).Decode(&account)
	if err != nil {
		log.Fatal(err.Error())
	}

	json.NewEncoder(res).Encode(account.SubAccount)
}
