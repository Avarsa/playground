package main

import (
	"log"
	"net/http"
)

//Running the user-authentication service
func main() {
	log.Fatal(http.ListenAndServe(":4000", http.HandlerFunc(UserAuthService)))
}
