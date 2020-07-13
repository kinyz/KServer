package msg

// Chat 聊天Key
const (
	ChatTopic          = "Chat_"
	ChatListeningTopic = "ChatListening_"
	ChatGroupTopic     = "ChatGroup_"

	ChatId      = 20000 // 聊天主线程id
	ChatWord    = 20001 // 发送世界
	ChatPrivate = 20002 // 发送私人
	ChatGroup   = 20003 //发送组别

	// Type
	ChatTypeText   = 1 // 文本
	ChatTypeVoice  = 2 //语音
	ChatTypeItem   = 3 // 物品
	ChatTypeImg    = 4 // 图片
	ChatTypeNotice = 5 // 公告
)
