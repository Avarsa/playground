package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/Avarsa/playground/tree/main/golang/simple-conditional-server/test_utils"
)

func TestUserService(t *testing.T) {

	//stub for test
	details := &Details{

		name: "user microservice",
		profile: Profile{
			Username: "Roger",
			Dob:      "08/08/1981",
			Age:      "39",
			Email:    "roger@goat.com",
			Phone:    "20",
		},
	}
	server := &UserService{
		details,
	}

	t.Run("inquiring the name", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/name", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		got := response.Body.String()
		want := "user microservice"
		test_utils.AssertString(t, got, want)

	})

	t.Run("Fetching user profiles", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/user/profile", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		got := response.Body.String()
		want := `{"Username":"Roger","Dob":"08/08/1981","Age":"39","Email":"roger@goat.com","Phone":"20"}`
		if !(reflect.DeepEqual(got, want)) {
			t.Errorf("got %s, want %s", got, want)
		}

	})
}
