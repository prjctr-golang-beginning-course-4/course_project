package model

import (
	"encoding/json"
	"net/http"
)

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUserLoginFromRequest(r *http.Request) (*UserLogin, error) {
	var newUserLogin UserLogin
	err := json.NewDecoder(r.Body).Decode(&newUserLogin)
	return &newUserLogin, err
}
