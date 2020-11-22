package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func allUsers(w http.ResponseWriter, r *http.Request) {
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
	var user UserWithPassword

	for _, u := range Users {
		if u.UID == user.UID || u.Name == user.Name || u.Email == user.Email {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "%v", "user already exists")
			return
		}
	}

	err := json.Unmarshal(reqBody, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return
	}
	actualUser := user.User
	actualUser.password = user.Password
	Users = append(Users, actualUser)
	json.NewEncoder(w).Encode(actualUser)
}

func addFriend(w http.ResponseWriter, r *http.Request) {
	//add the user posting to the friend list of the user who is in the URL
	vars := mux.Vars(r)
	friend := vars["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	var self string
	err := json.Unmarshal(reqBody, &self)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return
	}
	friendsList := FriendsMap[friend]
	friendsList = append(friendsList, self)
	FriendsMap[friend] = friendsList
	json.NewEncoder(w).Encode(FriendsMap[friend])
}

func showFriends(w http.ResponseWriter, r *http.Request) {
	//show friends of the user in the URL
	vars := mux.Vars(r)
	friend := vars["id"]
	if _, ok := FriendsMap[friend]; !ok {
		FriendsMap[friend] = []string{}
	}
	json.NewEncoder(w).Encode(FriendsMap[friend])
}
