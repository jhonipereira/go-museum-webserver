package api

import (
	"encoding/json"
	"jhonidev/go/goWebServer/data"
	"net/http"
	"strconv"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// r.URL.Query().Get("id")
	id := r.URL.Query()["id"]
	if id != nil { //try to search an individual record
		finalId, err := strconv.Atoi(id[0])
		if err == nil && finalId < len(data.GetAll()) {
			json.NewEncoder(w).Encode(data.GetAll()[finalId])
		} else {
			http.Error(w, "Invalid parameters", http.StatusBadRequest)
		}
	} else { //all records
		json.NewEncoder(w).Encode(data.GetAll())
	}
}
