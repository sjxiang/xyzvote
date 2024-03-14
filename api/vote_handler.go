package api

import (
	"xyzvote/db"
)


type VoteHandler struct {
	// db
	voteStore db.VoteStore
	userStore db.UserStore
}


func NewVoteHandler(voteStore db.VoteStore, userStore db.UserStore) *VoteHandler {
	return &VoteHandler{
		voteStore: voteStore,
		userStore: userStore,
	}
}

