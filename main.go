package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage")
}

func handleRequests(addr string) {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.Use(jsonContentType)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/users", returnAllUsers)
	myRouter.HandleFunc("/user", createNewUser).Methods("POST")
	myRouter.HandleFunc("/user/{id}", returnSingleUser)

	log.Fatal(http.ListenAndServe(addr, myRouter))
}

func jsonContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

type User struct {
	UID   string //`json:"Title"`
	Name  string
	Email string
}

var Users []User

func main() {
	var port string
	flag.StringVar(&port, "port", "10000", "Port of the interop server")
	flag.Parse()

	Users = []User{
		{UID: "1", Name: "testuser1", Email: "testuser@gmail.com"},
		{UID: "2", Name: "billy bob", Email: "billyb@gmail.com"},
	}
	handleRequests(fmt.Sprintf(":%v", port))
}

func returnAllUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Users)
}

func returnSingleUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, user := range Users {
		if user.UID == id {
			json.NewEncoder(w).Encode(user)
		}
	}
}

func createNewUser(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user User
	err := json.Unmarshal(reqBody, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return
	}
	Users = append(Users, user)
	json.NewEncoder(w).Encode(user)
}
