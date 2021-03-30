package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UserService struct {
	data DataStore
}

func (u *UserService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router := http.NewServeMux()
	router.Handle("/name", http.HandlerFunc(u.RelayMetaData))
	router.Handle("/user/profile", http.HandlerFunc(u.RelayUserProfile))
	router.ServeHTTP(w, r)
}

func (u *UserService) RelayUserProfile(w http.ResponseWriter, r *http.Request) {
	js, _ := json.Marshal(u.data.GetProfile())
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (u *UserService) RelayMetaData(w http.ResponseWriter, r *http.Request) {
	name := u.data.GetName()
	fmt.Fprintf(w, name)
}
