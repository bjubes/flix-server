package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Flix Server API")
}

func handleRequests(addr string) {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.Use(jsonContentType)

	myRouter.HandleFunc("/", home)
	myRouter.HandleFunc("/users", allUsers)
	myRouter.HandleFunc("/user/{id}/friend", addFriend).Methods("POST")
	myRouter.HandleFunc("/user/{id}/friends", showFriends)

	myRouter.HandleFunc("/user/{id}", returnSingleUser)

	myRouter.HandleFunc("/groups", allGroups)
	myRouter.HandleFunc("/group", createNewGroup).Methods("POST")
	myRouter.HandleFunc("/group/{id}/leave", leaveGroup).Methods("POST")
	myRouter.HandleFunc("/group/{id}/like", likeMovie).Methods("POST")
	myRouter.HandleFunc("/group/{id}", returnSingleGroup)

	myRouter.HandleFunc("/login", loginUser).Methods("POST")
	myRouter.HandleFunc("/register", createNewUser).Methods("POST")

	log.Fatal(http.ListenAndServe(addr, myRouter))
}

func jsonContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

var Users []User
var Groups []Group
var FriendsMap map[string][]string

func main() {
	var port int
	flag.IntVar(&port, "port", 10000, "Port of the flix server")
	flag.Parse()

	if p := os.Getenv("PORT"); p != "" {
		port, _ = strconv.Atoi(p)
	}
	FriendsMap = make(map[string][]string)
	if strings.ToLower(os.Getenv("PROD")) != "true" {
		Users = []User{
			{UID: "1", Name: "testuser1", Email: "testuser@gmail.com", password: "test"},
			{UID: "2", Name: "testuser2", Email: "billyb@gmail.com"},
		}
		votes := make(map[string][]string)
		votes["1"] = []string{}
		votes["2"] = []string{}
		Groups = []Group{
			NewGroup("1", "2"),
		}
	}
	handleRequests(fmt.Sprintf(":%v", port))
}
