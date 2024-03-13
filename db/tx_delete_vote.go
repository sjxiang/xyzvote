package db

import (
	"context"

	"gorm.io/gorm"
)

// 删除某项投票相关（设置、投票选项统计分析、用户投票记录）
func (s *MySQLVoteStore) DeleteVoteTx(ctx context.Context, voteId int64) error {

	if err := s.database.Transaction(func(tx *gorm.DB) error {
		// 1. 删除某项问卷调查
		if err := tx.Where("id = ?", voteId).Delete(&Vote{}).Error; err != nil {
			return err 
		}

		// 2. 删除相关选项
		if err := tx.Where("vote_id = ?", voteId).Delete(&Option{}).Error; err != nil {
			return err 
		}

		// 3. 删除相关投票记录
		if err := tx.Where("vote_id = ?", voteId).Delete(&VoteRecord{}).Error; err != nil {
			return err 
		}

		return nil 
	}); err != nil {
		return err 
	}

	return nil 
}

const (
	deleteVote1 = `
delete from vote where id = ? limit 1
`
	deleteVote2 = `
delete from option where vote_id = ?
`
	deleteVote3 = `
delete from vote_record where vote_id = ?
`

)
func (s *MySQLVoteStore) XDeleteVoteTx(ctx context.Context, voteId int64) error {

	if err := s.database.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(deleteVote1, voteId).Error; err != nil {
			return err 
		}
		if err := tx.Exec(deleteVote2, voteId).Error; err != nil {
			return err 
		}
		if err := tx.Exec(deleteVote3, voteId).Error; err != nil {
			return err 
		}
		
		return nil 
	}); err != nil {
		return err 
	}

	return nil 
}
