package repository

import (
	"context"
	"course/src/model"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
)

var ctx = context.Background()

type UserRepository struct {
	DB    *gorm.DB
	Cache *redis.Client
}

func InitUserRepository(db *gorm.DB, rdb *redis.Client) UserRepository {
	rep := UserRepository{DB: db, Cache: rdb}

	adm, _ := rep.GetUserByEmail(&model.User{Email: os.Getenv("ADMIN_EMAIL")})
	adm, _ = rep.StoreUser(adm)
	role, _ := rep.StoreRole(&model.Role{Name: os.Getenv("ADMIN_ROLE")})
	rep.AssignRoleToUser(adm, role)

	return rep
}

func (r *UserRepository) GetAllUsers() []model.User {
	var users []model.User

	if err := r.DB.Find(&users).Error; err != nil {
		return []model.User{}
	}

	return users
}

func (r *UserRepository) StoreUser(u *model.User) (*model.User, error) {
	if u.ID != 0 {
		return u, nil
	}

	existedUser, err := r.GetUserByEmail(u)
	if (err != nil && err != gorm.ErrRecordNotFound) || existedUser != nil {
		return nil, errors.New("email isn't valid")
	}

	existedUser, err = r.GetUserByLogin(u)
	if (err != nil && err != gorm.ErrRecordNotFound) || existedUser != nil {
		return nil, errors.New("login isn't valid")
	}

	u.Password = model.HashPassword(u.Password)
	result := r.DB.Create(u)
	if result.Error != nil {
		log.Fatal(result.Error)
		return nil, result.Error
	}

	return u, nil
}

func (r *UserRepository) StoreRole(role *model.Role) (*model.Role, error) {
	if role.ID != 0 {
		return role, nil
	}

	existedRole := &model.Role{}
	err := r.DB.Where("name = ?", role.Name).FirstOrCreate(existedRole, role).Error
	if err != nil {
		return nil, err
	}

	return existedRole, nil
}

func (r *UserRepository) GetUserByEmail(u *model.User) (*model.User, error) {
	existingUser := &model.User{}
	err := r.DB.Where("email = ?", u.Email).First(existingUser).Error
	if err != nil {
		return nil, err
	}

	return existingUser, err
}

func (r *UserRepository) GetUserById(id int) (*model.User, error) {
	u := &model.User{}
	err := r.DB.Where("id = ?", id).First(u).Error
	if err != nil {
		return nil, err
	}

	return u, err
}

func (r *UserRepository) GetUserByLogin(u *model.User) (*model.User, error) {
	existingUser := &model.User{}
	err := r.DB.Where("login = ?", u.Login).First(existingUser).Error
	if err != nil {
		return nil, err
	}

	return existingUser, err
}

func (r *UserRepository) GetUserByToken(token string) (*model.User, error) {
	key := "token:" + token
	result, err := r.Cache.Get(ctx, key).Result()
	id, err := strconv.Atoi(result)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%d", id)
	u, err := r.GetUserById(id)
	return u, err
}

func (r *UserRepository) AssignRoleToUser(u *model.User, role *model.Role) bool {
	if r.UserHasRole(u, role) {
		return true
	}

	userRole := &model.UserRole{
		UserId: u.ID,
		RoleId: role.ID,
	}

	result := r.DB.Create(userRole)
	if result.Error != nil {
		return false
	}
	return true
}

func (r *UserRepository) UserHasRole(u *model.User, role *model.Role) bool {
	var userRole model.UserRole
	err := r.DB.Joins("JOIN users ON user_roles.user_id = users.id").
		Joins("JOIN roles ON user_roles.role_id = roles.id").
		Where("users.id = ? AND roles.name = ?", u.ID, role.Name).
		First(&userRole).Error
	if err != nil {
		return false
	}

	return true
}

func (r *UserRepository) UserIsAdmin(u *model.User) bool {
	return r.UserHasRole(u, &model.Role{Name: os.Getenv("ADMIN_ROLE")})
}

func (r *UserRepository) UserAssignAdmin(u *model.User) bool {
	adminRole, _ := r.StoreRole(&model.Role{Name: os.Getenv("ADMIN_ROLE")})
	return r.AssignRoleToUser(u, adminRole)
}

func (r *UserRepository) DeleteUser(u *model.User) error {
	result := r.DB.Delete(&model.User{}, u.ID)

	if result.Error != nil {
		return fmt.Errorf("failed to delete user with id %d: %v", u.ID, result.Error)
	}

	if result.RowsAffected <= 0 {
		return fmt.Errorf("no user found with id %d", u.ID)
	}

	return nil
}
