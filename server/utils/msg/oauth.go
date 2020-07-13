package msg

// Oauth 验证服务器 100 - 200
const (
	OauthTopic = "Oauth_" // 验证服务器监听topic
	// 发送id
	OauthId           = 100 // 总id 发送
	OauthAccount      = 101 // 验证账号msgId
	OauthAccountClose = 102 // 账号下线

	// 返回id
	OauthAccountSuccess           = 102 // 客户端成功返回代码
	OauthAccountError             = 103 // 客户端验证失败返回代码
	OauthAccountOnlineError       = 104 // 客户端在线返回代码
	OauthAccountTokenError        = 105 // 客户端Token失效返回代码
	OauthAccountSystemError       = 106 // 系统失败返回代码
	OauthAccountAccountStateError = 107 // 客户端账号被封停
	OauthAccountNotFindError      = 108 // 客户端账号找不到账号
	OauthAccountConnIdError       = 109 // 客户端连接ID重复

	// 状态id
	OauthAccountOnline       = 200 // 客户端在线id
	OauthAccountAccountState = 0   // 账号状态 0为正常
)
