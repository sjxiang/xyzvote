package db

import (
	"time"
	"xyzvote/consts"
)

type User struct {
    Id                 int64     `json:"id,omitempty" gorm:"column:id;primary_key"`
    UserId             string    `json:"user_id"      gorm:"column:user_id"`
    Username           string    `json:"username"     gorm:"column:username"`
    EncryptedPassword  string    `json:"-"            gorm:"column:password"`
    Email              string    `json:"email"        gorm:"column:email"`
    Gender             int8      `json:"gender"       gorm:"column:gender"`
    CreatedAt          time.Time `json:"created_at"   gorm:"column:created_at"`
    UpdatedAt          time.Time `json:"updated_at"   gorm:"column:updated_at"`
}

func (u *User) TableName() string {
	return consts.UserTableName
}

// 问卷表单设置
type Vote struct {
	Id        int64     `json:"id,omitempty" gorm:"column:id;primary_key"`
    Title     string    `json:"title"        gorm:"column:title"`
    Type      int32     `json:"type"         gorm:"column:type"`
    Status    int32     `json:"status"       gorm:"column:status"`
    Duration  int64     `json:"duration"     gorm:"column:duration"`
    UserId    int64     `json:"user_id"      gorm:"column:user_id"`
	CreatedAt time.Time `json:"created_at"   gorm:"column:created_at"`
    UpdatedAt time.Time `json:"updated_at"   gorm:"column:updated_at"`
}

func (v *Vote) TableName() string {
    return consts.VoteTableName
}


// 问卷表单选项
type Option struct {
    Id        int64     `json:"id"           gorm:"column:id;primary_key"`
    Name      string    `json:"name"         gorm:"column:name"`
    Count     int32     `json:"count"        gorm:"column:count"`
    VoteId    int64     `json:"vote_id"      gorm:"column:vote_id"`
    CreatedAt time.Time `json:"created_at"   gorm:"column:created_at"`
    UpdatedAt time.Time `json:"updated_at"   gorm:"column:updated_at"`
}

func (v *Option) TableName() string {
    return consts.OptionTableName
}


// 用户投票记录
type VoteRecord struct {
	Id           int64     `json:"id"             gorm:"column:id;primary_key"`
    UserId       string    `json:"user_id"        gorm:"column:user_id"`
	VoteId       int64     `json:"vote_id"        gorm:"column:vote_id"`
	OptionId     int64     `json:"option_id"      gorm:"column:option_id"`
    CreatedAt    time.Time `json:"created_at"     gorm:"column:created_at"`
}

func (v *VoteRecord) TableName() string {
    return consts.VoteRecordTableName
}