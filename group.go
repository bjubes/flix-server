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
	var group VotelessGroup
	err := json.Unmarshal(reqBody, &group)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return
	}
	g := NewGroup(group.ID, group.Name, group.MaxTime, group.UserIDs...)
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
					if len(group.UserIDs) == 0 {
						Groups = append(Groups[:j], Groups[j+1:]...)
						return
					}
					delete(group.UserVotesMap, userid)
					Groups[j] = group
					return
				}
			}
		}
	}

}

type likeMoviePayload struct {
	UserID string
	Movies []string
}

func likeMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupID := vars["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	var likes likeMoviePayload
	err := json.Unmarshal(reqBody, &likes)
	if err != nil || likes.Movies == nil || likes.UserID == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return
	}
	for i, group := range Groups {
		if group.ID == groupID {
			newLikes := []string{}
			for _, l := range likes.Movies {
				var uniq = true
				for _, v := range group.UserVotesMap[likes.UserID] {
					if v == l {
						uniq = false
					}
				}
				if uniq {
					newLikes = append(newLikes, l)
				}
			}
			group.UserVotesMap[likes.UserID] = append(group.UserVotesMap[likes.UserID], newLikes...)
			Groups[i] = group
			return
		}
	}
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
