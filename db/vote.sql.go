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



