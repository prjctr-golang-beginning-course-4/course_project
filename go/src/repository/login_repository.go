package repository

import (
	"course/src/model"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"log"
)

type LoginRepository struct {
	DB    *gorm.DB
	Cache *redis.Client
}

func InitLoginRepository(db *gorm.DB, rdb *redis.Client) LoginRepository {
	return LoginRepository{DB: db, Cache: rdb}
}

func (r *LoginRepository) StoreLogin(l *model.Login) (*model.Login, error) {
	result := r.DB.Create(l)
	if result.Error != nil {
		log.Fatal(result.Error)
		return nil, result.Error
	}

	return l, nil
}

func (r *LoginRepository) GetUserFromLogin(ul *model.UserLogin) (*model.User, error) {
	u := &model.User{}
	err := r.DB.Where("email = ? AND password = ?", ul.Email, model.HashPassword(ul.Password)).First(u).Error
	if err != nil {
		return nil, err
	}

	return u, nil
}
