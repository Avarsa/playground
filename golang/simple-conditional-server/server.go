package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//implementing the proxy micro service
func LandingServer(writer http.ResponseWriter, request *http.Request) {
	//not the best way to do this.
	//A new instance of router is created for every request, which is ineffcient.
	router := http.NewServeMux()
	router.Handle("/microservice/name", http.HandlerFunc(RelayServiceName))
	router.Handle("/user/profile", http.HandlerFunc(RelayUserProfile))
	router.ServeHTTP(writer, request)
}

//RelayServiceName talks to the user microservice to get its name
func RelayServiceName(writer http.ResponseWriter, request *http.Request) {
	resp, _ := http.Get("http://localhost:7000/name")
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Fprintf(writer, string(body))
}

//RelayUserProfile talks to the user microservice to fetch user profile after receving
//StatusOK from the auth-microservice
func RelayUserProfile(writer http.ResponseWriter, request *http.Request) {
	username := request.Header.Get("username")
	if verifiedUser(username) {
		resp, _ := http.Get("http://localhost:7000/user/profile")
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Fprintf(writer, string(body))
		return
	}
	writer.WriteHeader(http.StatusUnauthorized)
	fmt.Fprintf(writer, "Nothing to see here")
}

//verifiedUser takes the header value and sends it to the user auth microservice for verification
func verifiedUser(username string) bool {
	if len(username) != 0 {
		client := &http.Client{}
		resp, _ := client.Get("http://localhost:4000/auth")
		req, _ := http.NewRequest("GET", "http://localhost:4000/auth", nil)
		req.Header.Set("username", username)
		resp, _ = client.Do(req)
		if (resp.StatusCode) == 200 {
			return true
		}
	}
	return false
}
