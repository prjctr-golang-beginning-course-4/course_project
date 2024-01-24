package controller

import (
	"context"
	"course/src/model"
	"course/src/repository"
	"course/src/response"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"net/http"
	"time"
)

var ctx = context.Background()

type LoginController struct {
	Repository repository.LoginRepository
}

func InitLoginController(db *gorm.DB, rdb *redis.Client) *LoginController {
	return &LoginController{Repository: repository.InitLoginRepository(db, rdb)}
}

func (c *LoginController) Login(w http.ResponseWriter, r *http.Request) {
	ul, err := model.CreateUserLoginFromRequest(r)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	u, err := c.Repository.GetUserFromLogin(ul)
	if err != nil || u == nil {
		http.Error(w, "Please register before", http.StatusBadRequest)
		return
	}

	token := model.GenerateToken(30)
	key := fmt.Sprintf("token:%s", token)
	err = c.Repository.Cache.Set(ctx, key, u.ID, time.Hour).Err()
	if err != nil {
		http.Error(w, "Cache issue", http.StatusInternalServerError)
		return
	}

	l := &model.Login{UserId: u.ID}
	l, err = c.Repository.StoreLogin(l)

	if err != nil || l == nil {
		http.Error(w, "Failed to store login", http.StatusBadRequest)
		return
	}

	response.Response{Message: "Success", Data: token, Code: http.StatusOK}.JsonResponse(w)
}
