package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/users", returnAllUsers)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

type User struct {
	UID   string //`json:"Title"`
	Name  string `json:"desc"`
	Email string `json:"content"`
}

var Users []User

func main() {
	Users = []User{
		{UID: "1", Name: "testuser1", Email: "testuser@gmail.com"},
		{UID: "2", Name: "billy bob", Email: "billyb@gmail.com"},
	}
	handleRequests()
}

func returnAllUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Users)
}
