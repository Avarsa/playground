package main

import (
	"log"
	"net/http"
)

func main() {

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

	log.Fatal(http.ListenAndServe(":7000", server))
}
