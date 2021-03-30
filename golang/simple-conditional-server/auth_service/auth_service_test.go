package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	test_utils "github.com/Avarsa/playground/tree/main/golang/simple-conditional-server/test_utils"
)

func TestUserAuthService(t *testing.T) {
	t.Run("without username header", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/auth", nil)
		response := httptest.NewRecorder()
		UserAuthService(response, request)
		got := response.Code
		want := http.StatusUnauthorized
		test_utils.AssertStatus(t, got, want)

	})

	t.Run("With username header", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/auth", nil)
		request.Header.Set("username", "chomsky")
		response := httptest.NewRecorder()
		UserAuthService(response, request)
		got := response.Code
		want := http.StatusOK
		test_utils.AssertStatus(t, got, want)
	})
}
