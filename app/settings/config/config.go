package config

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"simpus/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectDB is
func ConnectDB(res http.ResponseWriter, req *http.Request) (*mongo.Database, error) {

	// params := mux.Vars(req)
	// Username := params["user"]
	// Password := params["pass"]

	clientOptions := options.Client()
	// credential := options.Credential{
	// 	Username: "admin",
	// 	Password: "admin",
	// }
	clientOptions.ApplyURI("mongodb://localhost:27017")
	// clientOptions.ApplyURI("mongodb://mongodb").SetAuth(credential)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Connect(context.Background())
	if err != nil {
		return nil, err
	}

	return client.Database("simpus"), nil
}

// IndexDB is
func IndexDB(res http.ResponseWriter, req *http.Request) {
	db, err := db.MongoDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	listdb, err := db.Client().ListDatabaseNames(context.TODO(), bson.D{})
	json.NewEncoder(res).Encode(listdb)
}

// IndexCollection is
func IndexCollection(res http.ResponseWriter, req *http.Request) {
	db, err := db.MongoDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	listdb, err := db.ListCollectionNames(context.TODO(), bson.D{})
	json.NewEncoder(res).Encode(listdb)
}

// Store is
func Store(res http.ResponseWriter, req *http.Request) {
	db, err := db.MongoDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	db.CreateCollection(context.Background(), "test1")
}

// CreateUser is
func CreateUser(res http.ResponseWriter, req *http.Request) {
	// 	r := client.Database(dbName).RunCommand(context.Background(),bson.D{{"createUser", userName},
	//     {"pwd", pass}, {"roles", []bson.M{{"role": roleName,"db":roldeDB}}}})
	// if r.Err() != nil {
	//     panic(r.Err())
	// }
}

// GetUser is
func GetUser(res http.ResponseWriter, req *http.Request) {
	db, err := db.MongoDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	user := db.RunCommand(context.Background(), bson.A{"getUser"})
	json.NewEncoder(res).Encode(user)
}
