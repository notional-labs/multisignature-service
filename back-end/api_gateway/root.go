package api_gateway

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/notional-labs/multisignature-service/db"
)

var g_db *mongo.Database

func InitAPI(db *mongo.Database) {
	router := mux.NewRouter().StrictSlash(true)

	g_db = db

	router.HandleFunc("/get-random", getRandom)
	router.HandleFunc("/save-tx", saveTxToDB).Methods("POST")
	router.HandleFunc("/save-sig", saveSigToDB).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getRandom(w http.ResponseWriter, r *http.Request) {

}

func saveTxToDB(w http.ResponseWriter, r *http.Request) {

	// checking header type to make sure json
	if r.Header.Get("Content-Type") != "application/json" {
		msg := "Content-Type header is not application/json"
		http.Error(w, msg, http.StatusUnsupportedMediaType)
		return
	}

	var bodyJson db.Tx
	err := json.NewDecoder(r.Body).Decode(&bodyJson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	}

	err = bodyJson.UpdateOne()
	if err != nil {
		log.Fatal(err)
	}
}

func saveSigToDB(w http.ResponseWriter, r *http.Request) {

}
