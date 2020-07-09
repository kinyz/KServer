package services

import (
	"KServer/library/old/message"
	"KServer/server/utils"
)

type CharHandle struct {
	sendService Send
}

func NewCharHandle() *CharHandle {

	return &CharHandle{
		//ChatMessage: message.NewChatMessage(),
	}
}
func (c *CharHandle) PreHandle(msg messagereq) {
	//fmt.Println("我来了")
	switch msg.GetTopic() {
	case utils.ChatTopic:
		c.PreKeyHandle(msg)
	}

}
func (c *CharHandle) PreKeyHandle(msg messagereq) {
	switch msg.GetKey() {
	case utils.ChatSendKey:
		c.PreSendHandle(msg)
	}
}

func (c *CharHandle) PreSendHandle(msg messagereq) {

	switch msg.GetId() {
	case utils.ChatSendWord:
		c.sendService.SendWord(msg)
	case utils.ChatSendPrivate:
		c.sendService.SendPrivate(msg)
	}
}
