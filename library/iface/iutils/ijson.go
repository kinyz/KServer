package iutils

type IJson interface {
	Encode(table interface{}) []byte
	Decode(data []byte, table interface{}) error
}
