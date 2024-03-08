package types

import "time"

type Vote struct {
	Id        int64     `json:"id,omitempty" gorm:"column:id;primary_key"`
    Title     string    `json:"title"        gorm:"column:title"`
    Type      int32     `json:"type"         gorm:"column:type"`
    Status    int32     `json:"status"       gorm:"column:status"`
    Time      int64     `json:"time"         gorm:"column:time"`
    UserId    int64     `json:"user_id"      gorm:"column:user_id"`
	CreatedAt time.Time `json:"created_at"   gorm:"column:created_at"`
    UpdatedAt time.Time `json:"updated_at"   gorm:"column:updated_at"`
}

func (v *Vote) TableName() string {
    return "vote"
}

type GetVoteParams struct {
	ID int64 `form:"id" binding:"required,min=1"`
}


type VoteOption struct {
    Id        int64     `json:"id"           gorm:"column:id;primary_key"`
    Name      string    `json:"name"         gorm:"column:name"`
    Count     int32     `json:"count"        gorm:"column:count"`
    VoteId    int64     `json:"vote_id"      gorm:"column:vote_id"`
    CreatedAt time.Time `json:"created_at"   gorm:"column:created_at"`
    UpdatedAt time.Time `json:"updated_at"   gorm:"column:updated_at"`
}

func (v *VoteOption) TableName() string {
    return "vote_opt"
}

type VoteOptionByUser struct {
	Id           int64     `json:"id"             gorm:"column:id;primary_key"`
    UserId       int64     `json:"user_id"        gorm:"column:user_id"`
	VoteId       int64     `json:"vote_id"        gorm:"column:vote_id"`
	VoteOptionId int64     `json:"vote_option_id" gorm:"column:vote_option_id"`
    CreatedAt    time.Time `json:"created_at"     gorm:"column:created_at"`
}

func (v *VoteOptionByUser) TableName() string {
    return "vote_opt_user"
}

type VoteParams struct {
	VoteId       int64     `json:"vote_id"`
	VoteOptions  []int64   `json:"vote_options"`
}