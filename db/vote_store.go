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
	GetOptionsByVoteID(ctx context.Context, voteId int64) ([]*types.VoteOption, error)
	InsertUserVoteRecord(ctx context.Context, userId, voteId int64, options []int64) error
	InsertVoteAndOption(ctx context.Context, vote types.Vote, options []types.VoteOption) error
	UpdateVote(ctx context.Context, params types.UpdateVoteParams) error
	UpdateOption(ctx context.Context, params types.UpdateOptionParams) error
	DeleteVote(ctx context.Context, voteId int64) error
}

type MySQLVoteStore struct {
	storage *gorm.DB
}

func NewMySQLVoteStore(storage *gorm.DB) *MySQLVoteStore {
	return &MySQLVoteStore{
		storage:   storage,
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

// 查询问卷调查详情
func (s *MySQLVoteStore) GetVoteByID(ctx context.Context, id int64) (*types.Vote, error) {
	var (
		item types.Vote
		err  error
	)
	err = s.storage.Table("vote").Where("id = ?", id).First(&item).Error
	return &item, err
}

// 获取问卷调查选项详情
func (s *MySQLVoteStore) GetOptionsByVoteID(ctx context.Context, voteId int64) ([]*types.VoteOption, error) {
	var (
		items []*types.VoteOption
		err   error
	)
	err = s.storage.Table("vote_opt").Where("vote_id = ?", voteId).Find(&items).Error
	return items, err
}

// 新建用户投票纪录
func (s *MySQLVoteStore) InsertUserVoteRecord(ctx context.Context, userId, voteId int64, options []int64) error {

	if err := s.storage.Transaction(func(tx *gorm.DB) error {

		for _, val := range options {
			// 更新投票数
			if err := tx.Table("vote_opt").Where("id = ?", val).Update("count", gorm.Expr("count + ?", 1)).Error; err != nil {
				return err
			}
	
			// 关联
			var item = types.VoteOptionByUser{
				VoteId:       voteId,
				UserId:       userId,
				VoteOptionId: val,
				CreatedAt:    time.Now(),
			} 

			// 新建用户投票记录
			if err := tx.Table("vote_opt_user").Create(&item).Error; err != nil {
				return err
			}
		}
		
		return nil
	}); err != nil {
		return err 
	}

	return nil
}


/*

CQRS 
        将 db 的操作分为两类，即命令 command 与查询 query。
        命令则是会引起数据发生变化操作的总称，即新增、更新、删除
        查询则和字面意思一样，即不会对数据产生变化的操作，只是按照某些条件查找

	命令
		新增 insert
		更新 update
		删除 delete

			0
				成功
			1
				500 db 操作失败
				409 冲突
		
	查询 
		query
			0 
				成功
			1
				500 db 查询失败
				404 db 查询为空				

 */


// 新建问卷调查，标题和选项
func (s *MySQLVoteStore) InsertVoteAndOption(ctx context.Context, vote types.Vote, options []types.VoteOption) error {

	// 事务，匿名函数实现
	// 开启
	if err := s.storage.Transaction(func(tx *gorm.DB) error {
		
		// 新增标题
		if err := tx.Table("vote").Create(&vote).Error; err != nil {
			return err
		}

		// 新增选项
		for _, val := range options {
			// 关联
			val.VoteId = vote.Id
			if err := tx.Table("vote_opt").Create(&val).Error; err != nil {
				return err 
			}
		}

		// 提交
		return nil
	}); err != nil {
		// 回滚
		return err 
	}

	return nil
}



/*

	gorm update 缺陷

 */

func (s *MySQLVoteStore) UpdateVote(ctx context.Context, params types.UpdateVoteParams) error {
	return s.storage.Table("vote").Save(&params).Error 
}


func (s *MySQLVoteStore) UpdateOption(ctx context.Context, params types.UpdateOptionParams) error {
	return s.storage.Table("vote_opt").Save(&params).Error 
}

func (s *MySQLVoteStore) DeleteVote(ctx context.Context, voteId int64) error {
	if err := s.storage.Transaction(func(tx *gorm.DB) error {
		// 1. 删除某项问卷调查
		if err := tx.Delete(&types.Vote{}).Where("id = ?", voteId).Error; err != nil {
			return err 
		}

		// 2. 删除该项问卷调查选项
		if err := tx.Delete(&types.VoteOption{}).Where("vote_id = ?", voteId).Error; err != nil {
			return err 
		}

		// 3. 删除该项问卷调查投票记录
		if err := tx.Delete(&types.VoteOptionByUser{}).Where("vote_id = ?", voteId).Error; err != nil {
			return err 
		}

		return nil 
	}); err != nil {
		return err 
	}

	return nil 
}