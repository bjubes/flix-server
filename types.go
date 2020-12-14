package main

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

type VotelessGroup struct {
	ID      string
	Name    string
	UserIDs []string
	MaxTime int
}
type Group struct {
	VotelessGroup
	UserVotesMap map[string][]string //  UID -> [movieID]
}

func NewGroup(id, name string, userIDs ...string) Group {
	votes := make(map[string][]string)
	for _, uid := range userIDs {
		votes[uid] = []string{}
	}
	return Group{
		VotelessGroup{
			Name:    name,
			ID:      id,
			UserIDs: userIDs,
		},
		votes,
	}
}
