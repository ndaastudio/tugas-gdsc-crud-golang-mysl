package UserController

import (
	UserModel "crud-golang-mysql/models"
	"encoding/json"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user UserModel.Schema
	json.NewDecoder(r.Body).Decode(&user)

	err := UserModel.CreateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
