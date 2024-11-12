package controllers

import (
	"encoding/json"
	"main/models"
	"main/services"
	"net/http"

	"github.com/google/uuid"
)

var receipts = make(map[string]int)

func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt models.Receipt

	json.NewDecoder(r.Body).Decode(&receipt)

	points := services.CalculatePoints(receipt)
	id := uuid.New().String()

	receipts[id] = points

	response := map[string]string{"id": id}

	json.NewEncoder(w).Encode(response)
}
