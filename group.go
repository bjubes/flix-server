package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func allGroups(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Groups)
}

func returnSingleGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, group := range Groups {
		if group.ID == id {
			json.NewEncoder(w).Encode(group)
		}
	}
}

func createNewGroup(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var userIDs []string
	err := json.Unmarshal(reqBody, &userIDs)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return
	}
	g := NewGroup(userIDs...)
	Groups = append(Groups, g)
	json.NewEncoder(w).Encode(g)
}
