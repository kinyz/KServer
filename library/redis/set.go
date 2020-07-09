package redis

import (
	tool "KServer/library/utils"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"github.com/golang/protobuf/proto"
)

type SetValue struct {
	Conn redis.Conn
	Key  string
	tool.Protobuf
}

func (s *SetValue) ProtoBuf(value proto.Message) (reply interface{}, err error) {
	defer s.Conn.Close()
	return s.Conn.Do("SET", s.Key, s.Protobuf.Encode(value))
}
func (s *SetValue) String(value string) (reply interface{}, err error) {
	defer s.Conn.Close()
	return s.Conn.Do("SET", s.Key, value)
}
func (s *SetValue) Json(value interface{}) (reply interface{}, err error) {
	defer s.Conn.Close()
	b, _ := json.Marshal(&value)
	return s.Conn.Do("SET", s.Key, b)
}
func (s *SetValue) Bytes(value []byte) (reply interface{}, err error) {
	defer s.Conn.Close()
	return s.Conn.Do("SET", s.Key, value)
}
func (s *SetValue) Value(value interface{}) (reply interface{}, err error) {
	defer s.Conn.Close()
	return s.Conn.Do("SET", s.Key, value)
}
