package msg

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
