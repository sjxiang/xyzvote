package db

import (
	"context"
	"time"
	"xyzvote/consts"

	"gorm.io/gorm"
)


func (s *MySQLVoteStore) EndVoteTx(ctx context.Context) error {
	
	if err := s.database.Transaction(func(tx *gorm.DB) error {
		
		items := make([]*Vote, 0)
		
		// 筛选尚未结束的投票
		if err := tx.Table(consts.VoteTableName).Where("status = ?", 0).Find(&items).Error; err != nil {
			return err
		}

		nowAt := time.Now().Unix()
		for _, vote := range items {

			endAt := vote.Duration + vote.CreatedAt.Unix()
			
			// 判断是否到期
			if endAt <= nowAt {
				// 设置投票已结束
				if err := tx.Table(consts.VoteTableName).Where("id = ?", vote.Id).Update("status", 1).Error; err != nil {
					return err
				}
			}
		}
		
		return nil 
	}); err != nil {
		return err
	}

	return nil 
}


const (
	endVote1 = `
select * from vote where status = ?
`
	endVote2 = `
update vote set status = 1 where id = ? limit 1
`

)

func (s *MySQLVoteStore) XEndVoteTx(ctx context.Context) error {
	
	if err := s.database.Transaction(func(tx *gorm.DB) error {
		
		var (
			items []*Vote
		)
		if err := s.database.Raw(endVote1, 0).Scan(&items).Error; err != nil {
			return err
		}
		
		nowAt := time.Now().Unix()
		for _, vote := range items {

			endAt := vote.Duration + vote.CreatedAt.Unix()
			if endAt <= nowAt {
				if err := s.database.Exec(endVote2, vote.Id).Scan(&items).Error; err != nil {
					return err
				}
			}
		}
		
		return nil 
	}); err != nil {
		return err
	}

	return nil 
}




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
