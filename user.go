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
