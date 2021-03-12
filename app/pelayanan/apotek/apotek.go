package apotek

import (
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Apotek struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	NIK       int                `bson:"nik,omitempty"`
	Obat      string             `bson:"obat,omitempty"`
	CreatedAt primitive.DateTime `bson:"createdat,omitempty"`
}

func Index(res http.ResponseWriter, req *http.Request) {
	return
}

func Store(res http.ResponseWriter, req *http.Request) {
	return
}
