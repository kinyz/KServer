package services

import (
	"KServer/library/old/message"
	tool "KServer/library/utils"
	"fmt"
)

type Send struct {
	Tool tool.JsonTool
}

func (s *Send) SendWord(msg messagereq) {

	fmt.Println("SendWord", string(msg.GetData()))
}
func (s *Send) SendPrivate(msg messagereq) {
	fmt.Println("SendPrivate", string(msg.GetData()))

}
