package sdp

import (
	"errors"
	"strings"
)

// 连接数据，可以是会话级别，也可以是媒体级别
type Connection struct {
	NetworkType string // 网络类型
	AddressType string // 地址类型
	Address     string // 连接地址
}

func parseConnection(str string) (item Connection, err error) {
	item = Connection{}
	err = item.parse(str)
	return
}

// 字符串表达
func (c Connection) String() string {
	arr := []string{
		c.NetworkType,
		c.AddressType,
		c.Address,
	}
	return strings.Join(arr, seperator)
}

// 从文本中读取内容
func (c *Connection) parse(str string) (err error) {
	parts := strings.Split(str, seperator)
	if len(parts) != 3 {
		err = errors.New("sdp: connection is format error")
		return
	}
	c.NetworkType = parts[0]
	c.AddressType = parts[1]
	c.Address = parts[2]
	return
}
