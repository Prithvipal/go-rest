package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Prithvipal/go-rest/apierrors"
)

func writeSerivceError(w http.ResponseWriter, err error) {
	if err == apierrors.ErrNotFount {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func writeJSON(w http.ResponseWriter, records interface{}) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(records)
	if err != nil {
		log.Println("Error while getting TODO List", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Write(data)
}
