package sdp

import (
	"errors"
	"strings"
)

// 拥有者或者创建者和会话标识
type Basic struct {
	Username       string // 用户名
	SessionId      string // 会话ID
	SessionVersion string // 会话版本
	NetworkType    string // 网络类型
	AddressType    string // 地址类型
	UnicastAddress string // 任务创建计算机的地址
}

func parseBasic(str string) (item Basic, err error) {
	item = Basic{}
	err = item.parse(str)
	return
}

// 字符串表达
func (m Basic) String() string {
	arr := []string{
		m.Username,
		m.SessionId,
		m.SessionVersion,
		m.NetworkType,
		m.AddressType,
		m.UnicastAddress,
	}
	value := strings.Join(arr, seperator)
	return output(FieldBasic, value)
}

// 从文本中读取内容
func (m *Basic) parse(str string) (err error) {
	parts := strings.Split(str, seperator)
	if len(parts) != 6 {
		err = errors.New("sdp: basic is format error")
		return
	}
	m.Username = parts[0]
	m.SessionId = parts[1]
	m.SessionVersion = parts[2]
	m.NetworkType = parts[3]
	m.AddressType = parts[4]
	m.UnicastAddress = parts[5]
	return
}
