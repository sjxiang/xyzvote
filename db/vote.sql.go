package db

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"xyzvote/consts"
)


func (s *MySQLVoteStore) ListForms(ctx context.Context) ([]*Form, error) {
	items := make([]*Form, 0)
	if err := s.database.Table(consts.FormTableName).Find(&items).Error; err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, ErrRecordNoFound
	}

	return items, nil
}


type GetUserVoteRecordParams struct {
	UserId string
	FormId int64
}

func (s *MySQLVoteStore) GetUserVoteRecord(ctx context.Context, arg GetUserVoteRecordParams) ([]*VoteRecord, error) {
	items := make([]*VoteRecord, 0)
	if err := s.database.Table(consts.VoteRecordTableName).Where("user_id = ? and form_id = ?", arg.UserId, arg.FormId).Error; err != nil {
		return nil, err 
	}
	if len(items) == 0 {
		return nil, ErrRecordNoFound
	}

	return items, nil 
}



/*
	
	gorm update 缺陷

 */

 type UpdateFormParams struct {
    FormId    int64    
	Field     map[string]any
}

func (s *MySQLVoteStore) UpdateForm(ctx context.Context, arg UpdateFormParams) error {
	return s.database.Debug().Table(consts.FormTableName).Where("id = ?", arg.FormId).Save(arg.Field).Error 
}


func (s *MySQLVoteStore) GetForm(ctx context.Context, id int64) (*Form, error) {
	var item Form
	if err := s.database.Debug().Table(consts.FormTableName).Where("id = ?", id).First(&item).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNoFound
		} else {
			return nil, err
		}	
	}

	return &item, nil 
}

func (s *MySQLVoteStore) GetOptionByFormId(ctx context.Context, formId int64) ([]*Option, error) {
	items := make([]*Option, 0)
	if err := s.database.Debug().Table(consts.OptionTableName).Where("form_id = ?", formId).Find(&items).Error; err != nil {
		return nil, err 
	}
	if len(items) == 0 {
		return nil, ErrRecordNoFound
	}

	return items, nil 
}

/*

SELECT 
	* 
FROM
	`form` 
JOIN
	`form_opt` 
ON
	`form`.`id` = `form_opt`.`form_id` and `form`.`id`=6;

+----+-----------------+------+--------+----------+--------------------------------------+---------------------+---------------------+----+----------+------------+---------+---------------------+---------------------+
| id | title           | type | status | duration | user_id                              | created_at          | updated_at          | id | name     | vote_count | form_id | created_at          | updated_at          |
+----+-----------------+------+--------+----------+--------------------------------------+---------------------+---------------------+----+----------+------------+---------+---------------------+---------------------+
|  6 | today city walk |    0 |      0 |    86400 | f2d274f5-bb8b-4175-a60b-a2d113df4818 | 2024-03-14 08:26:14 | 2024-03-14 08:26:14 |  9 | nanjing  |          1 |       6 | 2024-03-14 08:26:14 | 2024-03-14 15:27:33 |
|  6 | today city walk |    0 |      0 |    86400 | f2d274f5-bb8b-4175-a60b-a2d113df4818 | 2024-03-14 08:26:14 | 2024-03-14 08:26:14 | 10 | shanghai |          1 |       6 | 2024-03-14 08:26:14 | 2024-03-14 15:27:33 |
+----+-----------------+------+--------+----------+--------------------------------------+---------------------+---------------------+----+----------+------------+---------+---------------------+---------------------+
*/