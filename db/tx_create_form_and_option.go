package db

import (
	"context"
	"time"
	"xyzvote/consts"

	"gorm.io/gorm"
)

/*

Tx

写操作（新增、修改、删除）
	冲突

读操作（查询）
	空

没有约束的话，先查后写

有约束的话，处理边界情况

 */

type CreateFormAndOptionTxParams struct {
	UserId      string
	Title       string
	Type        int32
	Status      int32
	Duration    int64
	OptionNames []string
	CreatedAt   time.Time 
}

// 创建表单和选项
func (s *MySQLVoteStore) CreateFormAndOptionTx(ctx context.Context, arg CreateFormAndOptionTxParams) error {
	if err := s.database.Debug().Transaction(func(tx *gorm.DB) error {
		var form Form

		form.UserId    = arg.UserId
		form.Title     = arg.Title
		form.Type      = arg.Type
		form.Status    = arg.Status
		form.Duration  = arg.Duration
		form.CreatedAt = arg.CreatedAt

		if err := tx.Table(consts.FormTableName).Create(&form).Error; err != nil {
			return err
		}
		
		for _, optionName := range arg.OptionNames {
			var option Option

			option.FormId    = form.Id
			option.Name      = optionName
			option.CreatedAt = arg.CreatedAt

			if err := tx.Table(consts.OptionTableName).Create(&option).Error; err != nil {
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

原生 SQL

INSERT INTO `form` 
	(`title`,`type`,`status`,`duration`,`user_id`,`created_at`,`updated_at`) 
VALUES 
	('today city walk', 0, 1, 86400, 'f2d274f5-bb8b-4175-a60b-a2d113df4818', NOW(), NOW());

INSERT INTO `form_opt` 
	(`name`,`vote_count`,`form_id`,`created_at`,`updated_at`) 
VALUES 
	('nanjing', 0, 4, NOW(), NOW());

INSERT INTO `form_opt` 
	(`name`,`vote_count`,`form_id`,`created_at`,`updated_at`)
VALUES 
	('shanghai', 0, 4, NOW(), NOW());

 */