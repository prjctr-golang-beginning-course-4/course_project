package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name string `gorm:"type:varchar(255);column:name;unique;not null" json:"name"`
}
