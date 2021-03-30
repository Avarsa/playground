package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	test_utils "github.com/Avarsa/playground/tree/main/golang/simple-conditional-server/test_utils"
)

func TestLandingServer(t *testing.T) {

	t.Run("inquiring the name of the service", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/microservice/name", nil)
		response := httptest.NewRecorder()
		LandingServer(response, request)
		got := response.Body.String()
		want := "user microservice"
		test_utils.AssertString(t, got, want)
	})

	t.Run("Fetching user profile without a username in the request header", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/user/profile", nil)
		response := httptest.NewRecorder()
		LandingServer(response, request)
		got := response.Code
		want := http.StatusUnauthorized
		test_utils.AssertStatus(t, got, want)

	})

	t.Run("Fetching user profiles with a username in the request header", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/user/profile", nil)
		request.Header.Add("username", "Jinish")
		response := httptest.NewRecorder()
		LandingServer(response, request)
		got := response.Code
		want := http.StatusOK
		test_utils.AssertStatus(t, got, want)

	})

}
