package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func loginUser(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var userLogin UserLogin
	err := json.Unmarshal(reqBody, &userLogin)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return
	}

	for _, user := range Users {
		if user.Name == userLogin.Username && user.password == userLogin.Password {
			json.NewEncoder(w).Encode(user)
			return
		}
	}

	//user was not found
	w.WriteHeader(http.StatusUnauthorized)
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var userLogin UserLogin
	err := json.Unmarshal(reqBody, &userLogin)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return
	}

	for _, user := range Users {
		if user.Name == userLogin.Username {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Username Taken")
			return
		}
	}

	//valid username, sign up user

}
