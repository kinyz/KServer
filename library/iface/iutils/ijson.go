package iutils

type IJson interface {
	StructToJson(tableStruct interface{}) (string, error)
	StructToByte(tableStruct interface{}) ([]byte, error)
	JsonToStruct(str string, tableStruct interface{}) error
}
