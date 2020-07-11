package pack

import (
	"KServer/library/iface/iutils"
	"KServer/library/utils"
)

type IToolPack interface {
	Json() iutils.IJson
	Byte() iutils.IByte
	Encrypt() iutils.IEncrypt
	Protobuf() iutils.IProtobuf
}

type ToolPack struct {
	IJson     iutils.IJson
	IByte     iutils.IByte
	IEncrypt  iutils.IEncrypt
	IProtobuf iutils.IProtobuf
}

func NewIToolPack() IToolPack {
	return &ToolPack{
		IJson:     utils.NewIJson(),
		IByte:     utils.NewIByte(),
		IEncrypt:  utils.NewIEncrypt(),
		IProtobuf: utils.NewIProtobuf(),
	}
}

func (t *ToolPack) Json() iutils.IJson {
	return t.IJson
}
func (t *ToolPack) Byte() iutils.IByte {
	return t.IByte
}
func (t *ToolPack) Encrypt() iutils.IEncrypt {
	return t.IEncrypt
}

func (t *ToolPack) Protobuf() iutils.IProtobuf {
	return t.IProtobuf
}
