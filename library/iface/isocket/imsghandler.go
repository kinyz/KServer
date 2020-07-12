package isocket

/*
	消息管理抽象层
*/
type IMsgHandle interface {
	DoMsgHandler(request IRequest)       // 马上以非阻塞方式处理消息
	AddHandle(id uint32, handle IHandle) // 为消息添加具体的处理逻辑
	AddCustomHandle(handle IHandle)      // 添加自定义头
	StartWorkerPool()                    // 启动worker工作池
	SendMsgToTaskQueue(request IRequest) // 将消息交给TaskQueue,由worker进行处理
}
