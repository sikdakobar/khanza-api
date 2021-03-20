package apotek

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"simpus/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Apotek struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	NIK       int                `bson:"nik,omitempty"`
	Obat      string             `bson:"obat,omitempty"`
	CreatedAt primitive.DateTime `bson:"createdat,omitempty"`
}

func Index(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	db, err := db.MongoDB()

	var apotek []Apotek

	// "date": time.Now().Format("1900-01-01")
	query, err := db.Collection("apotek").Find(context.Background(), bson.M{"date": "2021-03-11"})
	if err != nil {
		log.Fatal(err.Error())
	}

	defer query.Close(context.Background())

	for query.Next(context.Background()) {
		var data Apotek
		query.Decode(&data)
		apotek = append(apotek, data)
	}

	// for i := range apotek {
	// 	apotek[i].No_Urut = i + 1
	// }

	json.NewEncoder(res).Encode(apotek)
}

func Store(res http.ResponseWriter, req *http.Request) {
	return
}
