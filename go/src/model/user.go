package model

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
	"os"
)

type User struct {
	gorm.Model
	FirstName string `gorm:"type:varchar(255);column:first_name;not null" json:"firstName"`
	LastName  string `gorm:"type:varchar(255);column:last_name;not null" json:"lastName"`
	Email     string `gorm:"type:varchar(255);column:email;unique;not null" json:"email"`
	Login     string `gorm:"type:varchar(255);column:login;unique;not null" json:"login"`
	Password  string `gorm:"type:varchar(255);column:password;not null" json:"password"`
}

func CreateUserFromRequest(r *http.Request) (*User, error) {
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	return &newUser, err
}

func (u *User) IsAdmin() bool {
	return u.Email == "admin@company.com"
}

func (u *User) IsUser() bool {
	return u.Email != "admin@company.com"
}

func HashPassword(password string) string {
	concat := password + os.Getenv("SECRET")
	hashed := sha256.Sum256([]byte(concat))
	return hex.EncodeToString(hashed[:])
}

func CheckPassword(password string, hashed string) bool {
	return hashed == HashPassword(password)
}
