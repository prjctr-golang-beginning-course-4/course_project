package tests

import (
	"bytes"
	"course/src/model"
	"net/http/httptest"
	"testing"
)

func TestCreateUserLoginFromRequestPositive(t *testing.T) {
	newUserLogin := model.UserLogin{
		Email:    "test@gmail.com",
		Password: "test_password",
	}

	body := []byte("{\"Email\":\"test@gmail.com\",\"Password\":\"test_password\"}")
	req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(body))

	login, err := model.CreateUserLoginFromRequest(req)

	if err != nil {
		t.Fatalf("Didn't expect any errors, got: %v", err)
	}

	if login.Email != newUserLogin.Email || login.Password != newUserLogin.Password {
		t.Fatalf("Expected %s and %s, got: %v and %v ", newUserLogin.Email, newUserLogin.Password, login.Email, login.Password)
	}
}

func TestCreateUserLoginFromRequestNegative(t *testing.T) {
	newUserLogin := model.UserLogin{
		Email:    "test@gmail.com",
		Password: "test_password",
	}

	body := []byte("{\"Email\":\"test@gmail.de\",\"Password\":\"fail_password\"}")
	req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(body))

	login, _ := model.CreateUserLoginFromRequest(req)

	if login.Email == newUserLogin.Email || login.Password == newUserLogin.Password {
		t.Fatalf("Expected %s and %s, got: %v and %v ", newUserLogin.Email, newUserLogin.Password, login.Email, login.Password)
	}
}

func TestCreateUserLoginFromRequestDecodeNegative(t *testing.T) {
	body := []byte("sagddiogh")
	req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(body))

	_, err := model.CreateUserLoginFromRequest(req)

	if err == nil {
		t.Fatalf("Expected an error, didn't get one")
	}
}
