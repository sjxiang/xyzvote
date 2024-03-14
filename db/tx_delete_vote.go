package db

import (
	"context"

	"gorm.io/gorm"
)

// 清除某项问卷调查相关（表单设置、选项、用户投票记录）
func (s *MySQLVoteStore) DeleteVoteTx(ctx context.Context, formId int64) error {

	if err := s.database.Debug().Transaction(func(tx *gorm.DB) error {
		// 1. 删除表单
		if err := tx.Where("id = ?", formId).Delete(&Form{}).Error; err != nil {
			return err 
		}

		// 2. 删除选项
		if err := tx.Where("form_id = ?", formId).Delete(&Option{}).Error; err != nil {
			return err 
		}

		// 3. 删除投票记录
		if err := tx.Where("form_id = ?", formId).Delete(&VoteRecord{}).Error; err != nil {
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

	if err := s.database.Debug().Transaction(func(tx *gorm.DB) error {
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
