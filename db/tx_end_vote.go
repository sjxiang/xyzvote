package db

import (
	"context"
	"time"
	"xyzvote/consts"

	"gorm.io/gorm"
)


// 投票截止（表单设置）
func (s *MySQLVoteStore) EndVoteTx(ctx context.Context) error {
	
	if err := s.database.Debug().Transaction(func(tx *gorm.DB) error {
		
		items := make([]*Form, 0)
		
		// 筛选尚未结束的
		if err := tx.Table(consts.FormTableName).Where("status = ?", 0).Find(&items).Error; err != nil {
			return err
		}

		nowAt := time.Now().Unix()
		for _, vote := range items {

			endAt := vote.Duration + vote.CreatedAt.Unix()
			
			// 到期
			if endAt <= nowAt {
				// 设置表单状态
				if err := tx.Table(consts.FormTableName).Where("id = ?", vote.Id).Update("status", 1).Error; err != nil {
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

// 原生 raw sql 优化
func (s *MySQLVoteStore) XEndVoteTx(ctx context.Context) error {
	
	if err := s.database.Debug().Transaction(func(tx *gorm.DB) error {
		
		var (
			items []*Form
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



