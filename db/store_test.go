package db

import (
	"context"
	"testing"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func prepareVote(t *testing.T) VoteStore {
	var (
		dsn = "root:my-secret-pw@tcp(127.0.0.1:13306)/xyz_vote?parseTime=true"
	)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	return NewMySQLVoteStore(database)
}

func TestDoVote(t *testing.T) {

	voteStore := prepareVote(t)
	arg := DoVoteTxParams{
		UserId:    "f2d274f5-bb8b-4175-a60b-a2d113df4818",
		FormId:    6,
		OptIDs:    []int64{9, 10},
		CreatedAt: time.Now(),
	}

	err := voteStore.DoVoteTx(context.TODO(), arg)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(err)
}

func TestCreateFormAndOption(t *testing.T) {
	voteStore := prepareVote(t)
	arg := CreateFormAndOptionTxParams{
		UserId:      "f2d274f5-bb8b-4175-a60b-a2d113df4818",
		Title:       "today city walk",
		Type:        0,
		Status:      1,
		Duration:    86400,
		OptionNames: []string{"nanjing", "shanghai"},
		CreatedAt:   time.Now(),
	}

	err := voteStore.CreateFormAndOptionTx(context.TODO(), arg)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(err)
}


func TestDeleteVote(t *testing.T) {
	voteStore := prepareVote(t)
	err := voteStore.DeleteVoteTx(context.TODO(), 3)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(err)
}