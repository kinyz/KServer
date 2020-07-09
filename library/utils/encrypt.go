package utils

import (
	"KServer/library/iface/iutils"
	"crypto/md5"
	"encoding/hex"
	"github.com/satori/go.uuid"
	"math/rand"
	"time"
)

type Encrypt struct {
}

func NewIEncrypt() iutils.IEncrypt {
	return &Encrypt{}
}

// MD5生成
func (e *Encrypt) Md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// 新建随机token
func (e *Encrypt) NewToken() string {
	h := md5.New()
	h.Write([]byte(e.NewUuid()))
	return hex.EncodeToString(h.Sum(nil))
}

// 新建随机uuid4
func (e *Encrypt) NewUuid() string {
	// uuid.Must(uuid.NewV4())
	return uuid.NewV4().String()
}

// 取随机数字和字母
func (e *Encrypt) GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)

}
