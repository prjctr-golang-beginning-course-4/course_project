package controller

import (
	"course/src/model"
	"course/src/repository"
	"course/src/response"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"net/http"
)

type RegisterController struct {
	Repository repository.UserRepository
}

func InitRegisterController(db *gorm.DB, rdb *redis.Client) *RegisterController {
	return &RegisterController{Repository: repository.InitUserRepository(db, rdb)}
}

func (c *RegisterController) Register(w http.ResponseWriter, r *http.Request) {
	u, err := model.CreateUserFromRequest(r)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	u, err = c.Repository.StoreUser(u)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		return
	}

	response.Response{Message: "Success", Data: &u, Code: http.StatusCreated}.JsonResponse(w)
}
