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

func leaveGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reqBody, _ := ioutil.ReadAll(r.Body)
	groupID := vars["id"]
	var userID string
	err := json.Unmarshal(reqBody, &userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return
	}
	for j, group := range Groups {
		if group.ID == groupID {
			for i, userid := range group.UserIDs {
				if userid == userID {
					group.UserIDs = append(group.UserIDs[:i], group.UserIDs[i+1:]...)
					delete(group.UserVotesMap, userid)
					Groups[j] = group
					return
				}
			}
		}
	}

}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
