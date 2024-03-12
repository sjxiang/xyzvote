package types


type VoteRecord struct {

}

func (vr *VoteRecord) TableName() string {
    return "vote_opt_user"
}