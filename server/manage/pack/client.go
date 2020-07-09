package pack

import (
	"KServer/server/client"
	"KServer/server/manage/config"
	"KServer/server/utils/pd"
)

type IClientPack interface {
	AddClient(id uint32) bool
	GetOnlineNum() int
	GetClient(id uint32) client.IClient
	RemoveClient(id uint32)
	UpgradeClient(id uint32, account *pd.Account)
	GetClientByUUID(uuid string) client.IClient
}

type ClientPack struct {
	ConnId map[string]uint32
	Client map[uint32]client.IClient
}

func NewIClientPack(config *config.ManageConfig) IClientPack {
	if config.Client {
		return &ClientPack{
			Client: make(map[uint32]client.IClient),
			ConnId: make(map[string]uint32),
		}
	}
	return nil
}

func (m *ClientPack) AddClient(id uint32) bool {
	if m.Client[id] != nil {
		return false
	}

	m.Client[id] = &client.Client{
		Oauth:   false,
		Account: &pd.Account{},
	}

	return true
}
func (m *ClientPack) GetOnlineNum() int {
	return len(m.Client)
}

// 升级客户端
func (m *ClientPack) UpgradeClient(id uint32, account *pd.Account) {
	m.Client[id].SetOauth(true)
	m.Client[id].SetAccount(account)
	m.ConnId[m.Client[id].GetUUID()] = id

}
func (m *ClientPack) GetClientByUUID(uuid string) client.IClient {

	return m.Client[m.ConnId[uuid]]
}
func (m *ClientPack) GetClient(id uint32) client.IClient {
	return m.Client[id]
}
func (m *ClientPack) RemoveClient(id uint32) {

	if m.Client[id] != nil {
		delete(m.ConnId, m.Client[id].GetUUID())
		delete(m.Client, id)
	}

}
