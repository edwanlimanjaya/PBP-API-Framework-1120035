package controllers

import (
	"encoding/json"
	"exploration/model"
	"net/http"
)

func successResponseUsers(w http.ResponseWriter, users []model.User) {
	var response model.ResponseUsers

	response.Status = 200
	response.Message = "Success"
	response.Data = users

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func successResponse(w http.ResponseWriter) {
	var response model.Response

	response.Status = 200
	response.Message = "Success"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func rederactionResponse(w http.ResponseWriter) {
	var response model.Response

	response.Status = 305
	response.Message = "Not Modified"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func errorResponseClient(w http.ResponseWriter) {
	var response model.Response

	response.Status = 400
	response.Message = "Bad request"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func errorResponseServer(w http.ResponseWriter) {
	var response model.Response

	response.Status = 500
	response.Message = "Internal Server Error"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
