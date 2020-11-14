package main

import (
	"fmt"
	"math/rand"
)

type UserLogin struct {
	Username string
	Password string
}

type User struct {
	UID      string //`json:"uid"`
	Name     string
	Email    string
	password string //doesnt show in JSON
}

type UserWithPassword struct {
	User
	Password string
}

type Group struct {
	ID           string
	UserIDs      []string
	UserVotesMap map[string][]string //  UID -> [movieID]
}

func NewGroup(userIDs ...string) Group {
	votes := make(map[string][]string)
	for _, uid := range userIDs {
		votes[uid] = []string{}
	}
	id := fmt.Sprintf("%v", rand.Int())
	return Group{
		ID:           id,
		UserIDs:      userIDs,
		UserVotesMap: votes,
	}
}
