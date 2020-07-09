package utils

type IEncrypt interface {
	Md5V(str string) string
	NewToken() string
	NewUuid() string
	GetRandomString(l int) string
}
