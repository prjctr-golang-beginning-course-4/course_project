package controller

import (
	"course/src/model"
	"course/src/repository"
	"course/src/response"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
)

type UserController struct {
	Repository repository.UserRepository
}

func InitUserController(db *gorm.DB, rdb *redis.Client) *UserController {
	return &UserController{Repository: repository.InitUserRepository(db, rdb)}
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
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

func (c *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	response.Response{Message: "Success", Data: c.Repository.GetAllUsers(), Code: http.StatusOK}.JsonResponse(w)
}

func (c *UserController) Profile(w http.ResponseWriter, r *http.Request) {
	response.Response{Message: "Success", Data: c.getCurrentUser(r), Code: http.StatusOK}.JsonResponse(w)
}

func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.Atoi(params["userId"])
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	u, err := c.Repository.GetUserById(userId)
	if u == nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	response.Response{Message: "Success", Data: u, Code: http.StatusOK}.JsonResponse(w)
}

func (c *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.Atoi(params["userId"])
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	u, _ := c.Repository.GetUserById(userId)
	if u == nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	if u.ID == c.getCurrentUser(r).ID {
		http.Error(w, "can't delete yourself", http.StatusNotFound)
		return
	}

	err = c.Repository.DeleteUser(u)

	response.Response{Message: "Success", Data: err, Code: http.StatusOK}.JsonResponse(w)
}

func (c *UserController) AssignAdmin(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.Atoi(params["userId"])
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	u, _ := c.Repository.GetUserById(userId)
	if u == nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	c.Repository.UserAssignAdmin(u)

	response.Response{Message: "Success", Code: http.StatusOK}.JsonResponse(w)
}

func (c *UserController) getCurrentUser(r *http.Request) *model.User {
	header := r.Header.Get("Authorization")
	tokenParts := strings.Split(header, " ")
	u, _ := c.Repository.GetUserByToken(tokenParts[1])

	return u
}
