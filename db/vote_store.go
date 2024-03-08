package db

import (
	"context"
	"time"

	"gorm.io/gorm"

	"xyzvote/types"
)


type VoteStore interface {
	GetVotes(ctx context.Context) ([]*types.Vote, error)
	GetVoteByID(ctx context.Context, id int64) (*types.Vote, error) 
	GetOptionsByID(ctx context.Context, id int64) ([]*types.VoteOption, error)
	InsertVoteOptions(ctx context.Context, userId, voteId int64, options []int64) error
}

type MySQLVoteStore struct {
	storage *gorm.DB
	UserStore
}

func NewMySQLVoteStore(storage *gorm.DB, userStore UserStore) *MySQLVoteStore {
	return &MySQLVoteStore{
		storage: storage,
		UserStore: userStore,
	}
}

func (s *MySQLVoteStore) GetVotes(ctx context.Context) ([]*types.Vote, error) {
	var (
		items []*types.Vote
		err   error
	)
	err = s.storage.Table("vote").Find(&items).Error
	return items, err 
}

func (s *MySQLVoteStore) GetVoteByID(ctx context.Context, id int64) (*types.Vote, error) {
	var (
		item types.Vote
		err  error
	)
	err = s.storage.Table("vote").Where("id = ?", id).First(&item).Error
	return &item, err
}

func (s *MySQLVoteStore) GetOptionsByID(ctx context.Context, id int64) ([]*types.VoteOption, error) {
	var (
		items []*types.VoteOption
		err  error
	)
	err = s.storage.Table("vote_opt").Where("vote_id = ?", id).Find(&items).Error
	return items, err
}

func (s *MySQLVoteStore) InsertVoteOptions(ctx context.Context, userId, voteId int64, options []int64) error {

	for _, val := range options {
		var (
			err  error
			item types.VoteOptionByUser
		)
		err = s.storage.Table("vote_opt").Where("id = ?", val).Update("count", gorm.Expr("count + ?", 1)).Error
		if err != nil {
			return err
		}

		item.VoteId       = voteId
		item.UserId       = userId
		item.VoteOptionId = val
		item.CreatedAt    = time.Now()

		err = s.storage.Table("vote_opt_user").Create(&item).Error
		if err != nil {
			return err
		}
	}

	return nil
}