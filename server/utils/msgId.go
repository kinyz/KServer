package utils

// server key
const (
	AgentServerTopic      = "AgentServer_"
	AgentServerGroupTopic = "AgentServer_Group"
	AgentServerAllTopic   = "AgentServerAll_"

	AgentRegisterId = 10

	// 所有Agent服务器
	AgentAllServerId       = 1000
	AgentSendAllClientStop = 1001 // 所有客户端强制下线

	// 单服务器通知客户端
	AgentSendAllClient = 2001 // 通知所有客户端

)

// client key
const (
	ClientTopic          = "Client_"
	ClientTopicListening = "ChatMessage_Listening"
	ClientGroup          = "ChatServer"
	ClientSendKey        = "Send"

	ClientNotifyId = 5000 // 用于通知客户端的线程id
	ClientRemove   = 5001 // 用于通知客户端的线程id

	ClientSendWord    = 200
	ClientSendPrivate = 201

	ClientConnectID = 100 // 客户端数据验证

	ClientOnlineError = 103 // 客户端在线返回代码

	ClientConnectConnIdError = 108 // 客户端连接ID重复

	ClientOnline          = 200                  // 客户端在线id
	ClientConnectOauthKey = "ClientConnectOauth" // 客户端验证Key
	ClientLoginInfoKey    = "ClientLoginInfo_"   // 用户服务端Redis登陆头
	ClientAccountKey      = "Account"
	ClientUUIDKey         = "UUID"

	// 账号状态 0为正常
	ClientAccountState = 0
)

// Chat 聊天Key
const (
	ChatTopic          = "ChatMessage_test"
	ChatTopicListening = "ChatMessage_Listening"
	ChatGroup          = "ChatServer"
	ChatSendKey        = "ChatSend"
	ChatSendWord       = 200
	ChatSendPrivate    = 201

	ChatReceiveKey     = "ChatReceive"
	ChatReceiveWord    = 1200
	ChatReceivePrivate = 1201
)

// Oauth 验证服务器 100 - 200
const (
	OauthTopic = "Oauth_" // 验证服务器监听topic
	// 发送id
	OauthId           = 100 // 总id 发送
	OauthAccount      = 101 // 验证账号msgId
	OauthAccountClose = 102 // 验证账号msgId

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

// login 登陆服务器id
const (
	LoginTopic = "Login_"
)
