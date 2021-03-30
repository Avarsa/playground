package main

import (
	"net/http"
)

func UserAuthService(writer http.ResponseWriter, request *http.Request) {
	router := http.NewServeMux()
	router.Handle("/auth", http.HandlerFunc(CheckHeader))
	router.ServeHTTP(writer, request)
}

func CheckHeader(writer http.ResponseWriter, request *http.Request) {
	if len(request.Header.Values("username")) == 0 {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	writer.WriteHeader(http.StatusOK)

}
