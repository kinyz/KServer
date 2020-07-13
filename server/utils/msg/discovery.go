package msg

// 服务发现
const (
	ServiceDiscoveryTopic       = "Discovery_"       // 用于监听请求
	ServiceDiscoveryListenTopic = "DiscoveryListen_" // 用于监听服务的变化

	ServiceDiscoveryState = 200

	ServiceDiscoveryID = 10000 // 服务发现主通信ID
	// 以下MsgId
	ServiceDiscoveryRegister        = 10001 // 注册服务
	ServiceDiscoveryRegisterSuccess = 10002 // 注册服务成功回调
	ServiceDiscoveryRegisterFail    = 10003 // 注册服务失败回调
	ServiceDiscoveryLogoutService   = 10004 // 注销单个服务
	ServiceDiscoveryCloseService    = 10005 // 关闭总线程单个服务
	ServiceDiscoveryOpenService     = 10006 // 关闭总线程单个服务
	ServiceDiscoveryCheckAllService = 10007 // 请求所有已注册服务
	ServiceDiscoveryCheckService    = 10008 // 请求单个服务

)
