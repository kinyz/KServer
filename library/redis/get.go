package redis

import (
	tool "KServer/library/utils"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"github.com/golang/protobuf/proto"
)

type GetValue struct {
	Conn redis.Conn
	Key  string
	tool.Protobuf
}

func (g *GetValue) ProtoBuf(value proto.Message) error {
	bytes, err := redis.Bytes(g.Value())
	if err != nil {
		return err
	}
	return g.Protobuf.Decode(bytes, value)
}
func (g *GetValue) String() string {
	v, err := redis.String(g.Value())
	if err != nil {
		return ""
	}
	return v
}
func (g *GetValue) Json(value interface{}) error {
	bytes, _ := redis.Bytes(g.Value())
	err := json.Unmarshal(bytes, value)
	return err
}
func (g *GetValue) Bytes() []byte {
	v, err := redis.Bytes(g.Value())
	if err != nil {
		return []byte("")
	}
	return v
}
func (g *GetValue) Value() (reply interface{}, err error) {
	defer g.Conn.Close()
	return g.Conn.Do("Get", g.Key)
}
