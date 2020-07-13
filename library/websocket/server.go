package websocket

import (
	"KServer/library/iface/iwebsocket"
	"KServer/library/websocket/utils"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
)

//iServer 接口实现，定义一个Server服务类
type Server struct {
	//服务器的名称
	Name string
	//服务器协议 ws,wss
	Scheme string
	//服务绑定的IP地址
	IP string
	//服务绑定的端口
	Port int
	//协议
	Path string
	//当前Server的消息管理模块，用来绑定MsgId和对应的处理方法
	msgHandler iwebsocket.IMsgHandle
	//当前Server的链接管理器
	ConnMgr iwebsocket.IConnManager
	//该Server的连接创建时Hook函数
	OnConnStart func(conn iwebsocket.IConnection)
	//该Server的连接断开时的Hook函数
	OnConnStop func(conn iwebsocket.IConnection)
}

//连接信息
var upgrader = websocket.Upgrader{
	ReadBufferSize:  int(utils.GlobalObject.MaxPacketSize), //读取最大值
	WriteBufferSize: int(utils.GlobalObject.MaxPacketSize), //写最大值
	//解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

/*
  创建一个服务器句柄
*/
func NewWebsocket() iwebsocket.IServer {

	s := &Server{
		Name:       utils.GlobalObject.Name,
		Scheme:     utils.GlobalObject.Scheme,
		IP:         utils.GlobalObject.Host,
		Port:       utils.GlobalObject.TcpPort,
		Path:       utils.GlobalObject.Path, // 比如 /echo
		msgHandler: NewMsgHandle(),
		ConnMgr:    NewConnManager(),
	}
	return s
}

//============== 实现 iface.IServer 里的全部接口方法 ========

//websocket回调
func (s *Server) wsHandler(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("server wsHandler upgrade err:", err)
		return
	}
	// defer log.Println("server wsHandler client is closed")
	// defer conn.Close()

	// 判断是否超出个数
	if s.ConnMgr.Len() >= utils.GlobalObject.MaxConn {
		//todo 给用户发一个关闭连接消息
		log.Println("server wsHandler too many connection")

		conn.Close()
		return
	}

	log.Println("server wsHandler a new client coming ip:", conn.RemoteAddr())
	//处理新连接业务方法
	cid++

	dealConn := NewConntion(s, conn, cid, s.msgHandler)
	go dealConn.Start()
}

//全局conectionid 后续使用uuid生成
var cid uint32

//开启网络服务
func (s *Server) Start() {
	fmt.Printf("[START] Server name: %s,listenner at IP: %s, Port %d is starting\n", s.Name, s.IP, s.Port)

	//开启一个go去做服务端Linster业务
	go func() {
		//0 启动worker工作池机制
		s.msgHandler.StartWorkerPool()

		//已经监听成功
		fmt.Println("start ", s.Name, " succ, now listenning...")

		//TODO server.go 应该有一个自动生成ID的方法
		//	var cid uint32
		cid = 0

		http.HandleFunc("/"+s.Path, s.wsHandler)
		err := http.ListenAndServe(s.IP+":"+strconv.Itoa(int(s.Port)), nil)
		if err != nil {
			log.Println("server start listen error:", err)
		}
	}()
}

//停止服务
func (s *Server) Stop() {
	fmt.Println("[STOP] Zinx server , name ", s.Name)

	//将其他需要清理的连接信息或者其他信息 也要一并停止或者清理
	s.ConnMgr.ClearConn()
}

//运行服务
func (s *Server) Serve() {
	s.Start()

	//TODO Server.Serve() 是否在启动服务的时候 还要处理其他的事情呢 可以在这里添加

	//阻塞,否则主Go退出， listenner的go将会退出
	//select {}
}

//路由功能：给当前服务注册一个自定义头 找不到id时调用，供客户端链接处理使用
func (s *Server) AddCustomHandle(handle iwebsocket.IHandle) {
	s.msgHandler.AddCustomHandle(handle)
}

//路由功能：给当前服务注册一个路由业务方法，供客户端链接处理使用
func (s *Server) AddHandle(id uint32, handle iwebsocket.IHandle) {
	s.msgHandler.AddHandle(id, handle)
}

//得到链接管理
func (s *Server) GetConnMgr() iwebsocket.IConnManager {
	return s.ConnMgr
}

//设置该Server的连接创建时Hook函数
func (s *Server) SetOnConnStart(hookFunc func(iwebsocket.IConnection)) {
	s.OnConnStart = hookFunc
}

//设置该Server的连接断开时的Hook函数
func (s *Server) SetOnConnStop(hookFunc func(iwebsocket.IConnection)) {
	s.OnConnStop = hookFunc
}

//调用连接OnConnStart Hook函数
func (s *Server) CallOnConnStart(conn iwebsocket.IConnection) {
	if s.OnConnStart != nil {
		fmt.Println("---> CallOnConnStart....")
		s.OnConnStart(conn)
	}
}

//调用连接OnConnStop Hook函数
func (s *Server) CallOnConnStop(conn iwebsocket.IConnection) {
	if s.OnConnStop != nil {
		fmt.Println("---> CallOnConnStop....")
		s.OnConnStop(conn)
	}
}

func init() {
	//fmt.Println(zinxLogo)
	//fmt.Println(topLine)
	//fmt.Println(fmt.Sprintf("%s [Github] https://github.com/aceld                 %s", borderLine, borderLine))
	//fmt.Println(fmt.Sprintf("%s [tutorial] https://www.jianshu.com/p/23d07c0a28e5 %s", borderLine, borderLine))
	//fmt.Println(bottomLine)

	fmt.Printf("Socket Info MaxConn: %d, MaxPacketSize: %d\n",
		utils.GlobalObject.MaxConn,
		utils.GlobalObject.MaxPacketSize)
}
