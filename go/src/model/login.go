package model

import (
	"crypto/rand"
	"encoding/base64"
	"gorm.io/gorm"
)

type Login struct {
	gorm.Model
	UserId uint `gorm:"column:user_id;type:int;not null;foreignkey:ID;references:ID" json:"userId"`
	User   User `gorm:"foreignKey:UserId;references:ID"`
}

func GenerateToken(n int) string {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(b)
}
