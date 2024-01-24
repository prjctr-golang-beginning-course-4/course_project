package tests

import (
	"course/src/response"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJsonResponse(t *testing.T) {
	res := response.Response{
		Code: http.StatusOK,
	}

	newRecord := httptest.NewRecorder()
	res.JsonResponse(newRecord)
	result := newRecord.Result()

	if status := result.StatusCode; status != http.StatusOK {
		t.Errorf("Expected status code %v, got: %v", http.StatusOK, status)
	}

	if ctype := result.Header.Get("Content-Type"); ctype != "application/json" {
		t.Errorf("Expected content type 'application/json', got: %s", ctype)
	}

	expected := "{\"message\":\"\",\"data\":null,\"code\":200}\n"
	if newRecord.Body.String() != expected {
		t.Errorf("Unexpected body. Got: %s, want: %s", newRecord.Body.String(), expected)
	}
}
