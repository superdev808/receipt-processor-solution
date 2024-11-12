package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetPoints(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	points, exists := receipts[id]

	if !exists {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	response := map[string]int{"points": points}
	json.NewEncoder(w).Encode(response)
}
