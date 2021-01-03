package controller

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/Prithvipal/go-rest/apierrors"
	"github.com/Prithvipal/go-rest/models"
)

func writeSerivceError(w http.ResponseWriter, err error) {
	if _, ok := err.(*apierrors.NotFountErr); ok {
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

func getPageable(req *http.Request, defaultPage, defaultLimit int) (pageable models.Pageable, err error) {
	page := req.URL.Query().Get("page")
	if page == "" {
		pageable.Page = defaultPage
	} else {
		pageable.Page, err = strconv.Atoi(page)
		if err != nil || pageable.Page < 0 {
			return pageable, errors.New("Invalid page number")
		}
	}

	limit := req.URL.Query().Get("limit")

	if limit == "" {
		pageable.Limit = defaultLimit
	} else {
		pageable.Limit, err = strconv.Atoi(limit)
		if err != nil || pageable.Limit < 1 {
			return pageable, errors.New("Invalid limit")
		}
	}
	return
}
