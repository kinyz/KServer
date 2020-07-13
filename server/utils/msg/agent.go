package msg

// agent
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
