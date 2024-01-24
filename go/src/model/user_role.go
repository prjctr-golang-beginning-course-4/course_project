package model

import "gorm.io/gorm"

type UserRole struct {
	gorm.Model
	UserId uint `gorm:"column:user_id;not null" json:"userId"`
	RoleId uint `gorm:"column:role_id;not null" json:"roleId"`
	User   User `gorm:"foreignKey:UserId;references:ID"`
	Role   Role `gorm:"foreignKey:RoleId;references:ID"`
}
