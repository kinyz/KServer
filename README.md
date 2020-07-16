# KServer -- 碎片化服务框架 

kserver 组成部分 library manage

服务通信标准
topic     string    服务主题
id        uint32    路由主线程
msgid     uint32    路由执行线程
clientid  string    需要执行的clientid
serverid  string    转发的serverid
data      []byte    服务数据内容

目前已完成library
kafka       路由，通信数据解析
分布式lock   基本锁，时间锁，自动锁，次数锁，队列锁
redis       redis数据池
socket      修改zinx 
websocket   修改zinx
一些常用工具 protobuf json byte encrypt

目前已完成manage
discover 服务发现
socketclient websocketclient 客户端管理
message 碎片服务通信
kafka 路由管理 通信管理等
以及一些对库对封装

目前已完成基于kserver的分布式服务开发
socket/websocket 搭配服务发现模块来进行碎片服务转发 搭配client管理来做鉴权 
chat             聊天服务
lock             分布式锁的管理(目前需要用到kafka通信)
discover         配合socket和websocket进行服务发现 碎片服务热插拔
login            登陆服务获取socket和websocket所需要的鉴权信息
oauth            鉴权client的碎片服务

推荐使用rancher k8s管理系统来部署碎片服务 
具体文档及教程正在编写中，请先参考server里的使用 conf里有一些模版 可以自己定义

