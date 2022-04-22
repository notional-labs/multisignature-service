package api_gateway

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/notional-labs/multisignature-service/db"
)

func InitAPI() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/multisig/add", addMultisigPubkeyToKeyRing).Methods("POST")
	router.HandleFunc("/multisig/generate", generateMultisigPubkeyToKeyRing).Methods("POST")
	
	router.HandleFunc("/save/tx", saveTxToDB).Methods("POST")
	router.HandleFunc("/save/sig", saveSigToDB).Methods("POST")

	log.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func addMultisigPubkeyToKeyRing(w http.ResponseWriter, r *http.Request) {

}

func generateMultisigPubkeyToKeyRing(w http.ResponseWriter, r *http.Request) {

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
		return
	}

	err = bodyJson.UpdateOne()
	if err != nil {
		log.Fatal(err)
	}
}

func saveSigToDB(w http.ResponseWriter, r *http.Request) {
	// checking header type to make sure json
	if r.Header.Get("Content-Type") != "application/json" {
		msg := "Content-Type header is not application/json"
		http.Error(w, msg, http.StatusUnsupportedMediaType)
		return
	}

	var bodyJson db.Sign
	err := json.NewDecoder(r.Body).Decode(&bodyJson)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	err = bodyJson.UpdateOne()
	if err != nil {
		log.Fatal(err)
	}
}
