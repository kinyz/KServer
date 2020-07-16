package socket

import (
	"KServer/library/kiface/isocket"
	"KServer/library/socket/utils"
	"fmt"
	"strconv"
)

type MsgHandle struct {
	Handle         map[uint32]isocket.IHandle //存放每个Id 所对应的处理方法的map属性
	WorkerPoolSize uint32                     //业务工作Worker池的数量
	TaskQueue      []chan isocket.IRequest    //Worker负责取任务的消息队列
	CustomHandle   isocket.IHandle
	//Response map[uint32]ziface.IResponse
}

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Handle:         make(map[uint32]isocket.IHandle),
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize,
		//一个worker对应一个queue
		TaskQueue: make([]chan isocket.IRequest, utils.GlobalObject.WorkerPoolSize),
		//Response: make(map[uint32]ziface.IResponse),
	}
}

//将消息交给TaskQueue,由worker进行处理
func (mh *MsgHandle) SendMsgToTaskQueue(request isocket.IRequest) {
	//根据ConnID来分配当前的连接应该由哪个worker负责处理
	//轮询的平均分配法则

	//得到需要处理此条连接的workerID
	workerID := request.GetConnection().GetConnID() % mh.WorkerPoolSize
	//fmt.Println("Add ConnID=", request.GetConnection().GetConnID()," request msgID=", request.GetMsgID(), "to workerID=", workerID)
	//将请求消息发送给任务队列
	mh.TaskQueue[workerID] <- request
}

//马上以非阻塞方式处理消息
func (mh *MsgHandle) DoMsgHandler(request isocket.IRequest) {
	handler, ok := mh.Handle[request.GetMessage().GetId()]
	if !ok {
		//fmt.Println("Agreement Id = ", request.GetId(), " is not FOUND!")

		//如果没有id的请求将会执行自定义协议头
		if mh.CustomHandle != nil {
			fmt.Println("执行自定义头")
			mh.CustomHandle.PreHandle(request)
			mh.CustomHandle.PostHandle(request)
			return
		}

		request.GetConnection().SendBuffMsg([]byte("无服务"))
		request.GetConnection().Stop()
		return
	}

	//执行对应处理方法
	handler.PreHandle(request)
	//handler.RunMsg(request.GetMsgID())
	handler.PostHandle(request)
}

//添加一个自定义协议头 如果没有id的请求将会执行
func (mh *MsgHandle) AddCustomHandle(handle isocket.IHandle) {

	mh.CustomHandle = handle
	fmt.Println("Socket Add CustomHandle ")

}

//添加一个协议头
func (mh *MsgHandle) AddHandle(id uint32, handle isocket.IHandle) {
	//1 判断当前msg绑定的API处理方法是否已经存在
	if _, ok := mh.Handle[id]; ok {
		panic("repeated Agreement , Id = " + strconv.Itoa(int(id)))
	}
	//2 添加msg与api的绑定关系
	mh.Handle[id] = handle
	fmt.Println("Socket Add Handle = ", id)

}

//启动一个Worker工作流程
func (mh *MsgHandle) StartOneWorker(workerID int, taskQueue chan isocket.IRequest) {
	fmt.Println("Worker ID = ", workerID, " is started.")
	//不断的等待队列中的消息
	for {
		select {
		//有消息则取出队列的Request，并执行绑定的业务方法
		case request := <-taskQueue:
			mh.DoMsgHandler(request)
		}
	}
}

//启动worker工作池
func (mh *MsgHandle) StartWorkerPool() {
	//遍历需要启动worker的数量，依此启动
	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		//一个worker被启动
		//给当前worker对应的任务队列开辟空间
		mh.TaskQueue[i] = make(chan isocket.IRequest, utils.GlobalObject.MaxWorkerTaskLen)
		//启动当前Worker，阻塞的等待对应的任务队列是否有消息传递进来
		go mh.StartOneWorker(i, mh.TaskQueue[i])
	}
}
