package db

import (
	"context"
	"time"

	"gorm.io/gorm"

	"xyzvote/consts"
)


type DoVoteTxParams struct {
	UserId    string
	FormId    int64
	OptIDs    []int64
	CreatedAt time.Time
}


// 投票（查重、选项统计、用户投票记录）
func (s *MySQLVoteStore) DoVoteTx(ctx context.Context, arg DoVoteTxParams) error {

	if err := s.database.Debug().Transaction(func(tx *gorm.DB) error {

		// 是否有这项问卷表单
		forms := make([]*Form, 0)
		if err := tx.Table(consts.FormTableName).Where("id = ?", arg.FormId).Find(&forms).Error; err != nil {
			return err
		}
		if len(forms) == 0 {  
			return ErrRecordNoFound  // 没有这个投票
		}

		// 是否投过票
		history := make([]*VoteRecord, 0)
		if err := tx.Table(consts.VoteRecordTableName).Where("form_id = ? and user_id = ?", arg.FormId, arg.UserId).
			Find(&history).Error; err != nil {
			return err
		}
		if len(history) > 0 {  
			return ErrDuplicateVoteRecord  // 已投
		}

		for _, optID := range arg.OptIDs {
			// 选项投票数目
			if err := tx.Table(consts.OptionTableName).Where("id = ?", optID).Update("vote_count", gorm.Expr("vote_count + ?", 1)).Error; err != nil {
				return err
			}
	
			var item VoteRecord
			
			item.FormId    = arg.FormId
			item.OptionId  = optID
			item.UserId    = arg.UserId
			item.CreatedAt = arg.CreatedAt

			// 投票记录
			if err := tx.Table(consts.VoteRecordTableName).Create(&item).Error; err != nil {
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

SELECT 
	* 
FROM 
	`form_opt_user` 
WHERE 
	form_id = 1 and user_id = 'f2d274f5-bb8b-4175-a60b-a2d113df4818';


UPDATE 
	`form_opt` 
SET 
	`vote_count`=vote_count + 1 
WHERE 
	id = 1;

INSERT INTO `form_opt_user` 
	(`user_id`,`form_id`,`option_id`,`created_at`) 
VALUES 
	('f2d274f5-bb8b-4175-a60b-a2d113df4818', 1, 1, NOW());

*/

func (s *MySQLVoteStore) XDoVoteTx(ctx context.Context, arg DoVoteTxParams) error {
	if err := s.database.Debug().Transaction(func(tx *gorm.DB) error {

		var form Form
		if err := tx.Raw("select * from form where id = ?", arg.FormId).Scan(&form).Error; err != nil {
			return err
		}
		if form.Id == 0 {
			return ErrRecordNoFound
		}

		history := make([]*VoteRecord, 0)
		if err := tx.Raw("select * from form_opt_user where form_id = ? and user_id = ?", arg.FormId, arg.UserId).Scan(&history).Error; err != nil {
			return err
		}
		if len(history) > 0 {  
			return ErrDuplicateVoteRecord
		}

		for _, optID := range arg.OptIDs {
			
			if err := tx.Exec("update form_opt set count = count+1 where id = ? limit 1", optID).Error; err != nil {
				return err
			}
	
			var item VoteRecord
			
			item.FormId    = arg.FormId
			item.OptionId  = optID
			item.UserId    = arg.UserId
			item.CreatedAt = arg.CreatedAt

			// 投票记录
			if err := tx.Table(consts.VoteRecordTableName).Create(&item).Error; err != nil {
				return err
			}
		}
		
		return nil
	}); err != nil {
		return err 
	}

	return nil 
}