package types



type Option struct {

}

func (o *Option) TableName() string {
    return "vote_opt"
}