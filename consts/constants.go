package consts


const (
	HTTPServerAddres           = ":8090"
	
	UserTableName              = "user"
	FormTableName              = "form"
	OptionTableName            = "form_opt"
	VoteRecordTableName        = "form_opt_user"

	MySQLDefaultDSN            = "root:my-secret-pw@tcp(127.0.0.1:13306)/xyz_vote?charset=utf8&parseTime=True&loc=Local"
	
	RedisServerAddr            = "127.0.0.1:16379"
	RedisPassword              = ""
	RedisDatabaseNum           = 0

	CaptchaHeight              = 80    // 生成图片高度
	CaptchaWidth               = 240   // 生成图片宽度
	CaptchaLength              = 6     // 验证码的长度
	CaptchaMaxSkew             = 0.7   // 数字的最大倾斜角，越大，倾斜越狠，越不容易看懂
	CaptchaDotCount            = 80    // 图片背景里的混淆点数量
)

