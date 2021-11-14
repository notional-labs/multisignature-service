package api_gateway

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

var g_db *mongo.Database

func InitAPI(db *mongo.Database) {
	router := mux.NewRouter().StrictSlash(true)

	g_db = db

	router.HandleFunc("/save-tx", saveTxToDB).Methods("POST")
	router.HandleFunc("/save-sig", saveSigToDB).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func saveTxToDB(w http.ResponseWriter, r *http.Request) {

}

func saveSigToDB(w http.ResponseWriter, r *http.Request) {

}
