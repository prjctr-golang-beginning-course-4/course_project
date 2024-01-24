package tests

import (
	"course/src/model"
	"encoding/base64"
	"testing"
)

func TestGenerateTokenPositive(t *testing.T) {
	length := 10
	token := model.GenerateToken(length)

	decodedToken, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		t.Errorf("Failed to decode token")
	}

	if got, want := len(decodedToken), length; got != want {
		t.Errorf("Unexpected token length. Got: %d, want: %d", got, want)
	}

}

func TestGenerateTokenNegative(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Token generator didn't panic")
		}
	}()

	model.GenerateToken(-10)
}
