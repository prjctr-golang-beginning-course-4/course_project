package tests

import (
	"course/src/model"
	"course/src/repository"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

func TestGetUserByEmailPositive(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Couldn't create sqlmock: %s", err)
	}
	defer db.Close()

	u := &model.User{
		Email: "test@mail.com",
	}

	userRepository := &repository.UserRepository{DB: db}

	expectedUser := &model.User{
		Id:        1,
		FirstName: "Test",
		LastName:  "User",
		Email:     "test@mail.com",
		Login:     "testUser",
		Password:  "testPassword",
	}

	rows := sqlmock.NewRows([]string{"Id", "FirstName", "LastName", "Email", "Login", "Password"}).
		AddRow(expectedUser.Id, expectedUser.FirstName, expectedUser.LastName, expectedUser.Email, expectedUser.Login, expectedUser.Password)

	mock.ExpectQuery("^SELECT \\* FROM users WHERE email = \\?$").WithArgs(u.Email).WillReturnRows(rows)

	user, err := userRepository.GetUserByEmail(u)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if user.Id != expectedUser.Id {
		t.Errorf("Expected ID: %d, but got: %d", expectedUser.Id, user.Id)
	}
	if user.FirstName != expectedUser.FirstName {
		t.Errorf("Expected FirstName: %s, but got: %s", expectedUser.FirstName, user.FirstName)
	}
	if user.LastName != expectedUser.LastName {
		t.Errorf("Expected LastName: %s, but got: %s", expectedUser.LastName, user.LastName)
	}
	if user.Email != expectedUser.Email {
		t.Errorf("Expected Email: %s, but got: %s", expectedUser.Email, user.Email)
	}
	if user.Login != expectedUser.Login {
		t.Errorf("Expected Login: %s, but got: %s", expectedUser.Login, user.Login)
	}
	if user.Password != expectedUser.Password {
		t.Errorf("Expected Password: %s, but got: %s", expectedUser.Password, user.Password)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestGetUserByEmailNegative(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Couldn't create sqlmock: %s", err)
	}
	defer db.Close()

	u := &model.User{
		Email: "test@mail.com",
	}

	userRepository := &repository.UserRepository{DB: db}
	mock.ExpectQuery("^SELECT \\* FROM users WHERE id = \\?$").WithArgs(u.Id).WillReturnError(errors.New("sql: no rows in result set"))

	user, err := userRepository.GetUserByEmail(u)

	if err == nil {
		t.Errorf("Expected error recieve none")
	}

	if user != nil {
		t.Errorf("Expected no user, get one")
	}
}

func TestGetUserByIdPositive(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Couldn't create sqlmock: %s", err)
	}
	defer db.Close()

	u := &model.User{
		Id: 1,
	}

	userRepository := &repository.UserRepository{DB: db}

	expectedUser := &model.User{
		Id:        1,
		FirstName: "Test",
		LastName:  "User",
		Email:     "test@mail.com",
		Login:     "testUser",
		Password:  "testPassword",
	}

	rows := sqlmock.NewRows([]string{"Id", "FirstName", "LastName", "Email", "Login", "Password"}).
		AddRow(expectedUser.Id, expectedUser.FirstName, expectedUser.LastName, expectedUser.Email, expectedUser.Login, expectedUser.Password)

	mock.ExpectQuery("^SELECT \\* FROM users WHERE id = \\?$").WithArgs(u.Id).WillReturnRows(rows)

	user, err := userRepository.GetUserById(u.Id)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if user.Id != expectedUser.Id {
		t.Errorf("Expected Id: %d, but got: %d", expectedUser.Id, user.Id)
	}
	if user.FirstName != expectedUser.FirstName {
		t.Errorf("Expected FirstName: %s, but got: %s", expectedUser.FirstName, user.FirstName)
	}
	if user.LastName != expectedUser.LastName {
		t.Errorf("Expected LastName: %s, but got: %s", expectedUser.LastName, user.LastName)
	}
	if user.Email != expectedUser.Email {
		t.Errorf("Expected Email: %s, but got: %s", expectedUser.Email, user.Email)
	}
	if user.Login != expectedUser.Login {
		t.Errorf("Expected Login: %s, but got: %s", expectedUser.Login, user.Login)
	}
	if user.Password != expectedUser.Password {
		t.Errorf("Expected Password: %s, but got: %s", expectedUser.Password, user.Password)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestGetUserByIdNegative(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Couldn't create sqlmock: %s", err)
	}
	defer db.Close()

	u := &model.User{
		Id: 1,
	}

	userRpository := &repository.UserRepository{DB: db}
	rows := sqlmock.NewRows([]string{"Id", "FirstName", "LastName", "Email", "Login", "Password"})
	mock.ExpectQuery("^SELECT \\* FROM users WHERE id = \\?$").WithArgs(u.Id).WillReturnRows(rows)

	user, err := userRpository.GetUserById(u.Id)
	if err != nil {
		t.Errorf("Expected error recieve none")
	}

	if user != nil {
		t.Errorf("Expected no user, get one")
	}
}
