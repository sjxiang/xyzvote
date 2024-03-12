package db

import (
	"context"
	"time"
	"xyzvote/consts"

	"gorm.io/gorm"
)


type DoVoteTxParams struct {
	UserId    string
	VoteId    int64
	Options   []int64
	CreatedAt time.Time
}

type DoVoteTxResult struct {

}


// 参加投票活动（选项统计分析、用户投票记录）
func (s *MySQLVoteStore) DoVote(ctx context.Context, arg DoVoteTxParams) error {

	if err := s.database.Transaction(func(tx *gorm.DB) error {

		for _, val := range arg.Options {
			// 更新投票数
			if err := tx.Table(consts.OptionTableName).Where("id = ?", val).Update("count", gorm.Expr("count + ?", 1)).Error; err != nil {
				return err
			}
	
			// 关联
			var item VoteRecord
			
			item.VoteId    = arg.VoteId
			item.OptionId  = val
			item.UserId    = arg.UserId
			item.CreatedAt = arg.CreatedAt

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

