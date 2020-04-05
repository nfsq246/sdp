package sdp

import (
	"fmt"
	"strings"
)

// 完整的SDP消息
type Message struct {
	Version        string      // v 版本号
	Basic          Basic       // o 创建者/接收者和会话标识
	SessionName    string      // s 会话名称
	Connection     *Connection // c* 连接信息
	UnsupportLines []string    // 不支持处理的行
	Media          []Media     // 媒体信息
}

func NewMessage(str string) (item Message, err error) {
	item = Message{}
	err = item.parse(str)
	return
}

// 更新指定媒体的服务器地址和端口号
func (m *Message) UpdateMediaAddress(mediaType string, newDomain string, newPort int) {
	for i, _ := range m.Media {
		if m.Media[i].Info.MediaType == mediaType {
			m.Media[i].Info.Port = fmt.Sprintf("%d", newPort)
			if m.Media[i].Connection != nil {
				m.Media[i].Connection.Address = newDomain
			} else {
				m.Connection.Address = newDomain
			}
			return
		}
	}
}

// 字符串表达
func (m Message) String() (result string) {
	result += output(FieldVersion, m.Version)
	result += m.Basic.String()
	result += output(FieldSessionName, m.SessionName)
	if m.Connection != nil {
		result += output(FieldConnection, m.Connection.String())
	}
	for _, line := range m.UnsupportLines {
		result += line + CRLF
	}
	for _, media := range m.Media {
		result += media.String()
	}
	return
}

// 解析内容到消息中
func (m *Message) parse(str string) (err error) {
	isMedia := false
	for {
		// 已经找完了
		if len(str) == 0 {
			return
		}
		// 寻找下一个CRLF
		var content string
		loc := strings.Index(str, CRLF)
		if loc >= 0 {
			content = str[0:loc]
			str = str[loc+len(CRLF):]
		} else {
			content = str
			str = ""
		}
		// 空行跳过
		if len(content) == 0 {
			continue
		}
		// 解析行的key和value
		key, value, e := parse(content)
		if e != nil {
			err = e
			return
		}
		// 根据key处理value
		if isMedia {
			switch key {
			case FieldMediaInfo:
				if media, e := newMedia(value); e != nil {
					err = nil
				} else {
					m.Media = append(m.Media, media)
				}
			case FieldMediaTitle:
				m.Media[len(m.Media)-1].Title = value
			case FieldMediaConnection:
				if connection, e := parseConnection(value); e != nil {
					err = e
				} else {
					m.Media[len(m.Media)-1].Connection = &connection
				}
			case FieldMediaBandwidth, FieldMediaEncryptionKey, FieldMediaAttribute:
				m.Media[len(m.Media)-1].addUnsupportLine(content)
			}
		} else {
			switch key {
			case FieldVersion:
				m.Version = value
			case FieldBasic:
				m.Basic, err = parseBasic(value)
			case FieldSessionName:
				m.SessionName = value
			case FieldConnection:
				if connection, e := parseConnection(value); e != nil {
					err = e
				} else {
					m.Connection = &connection
				}
			case FieldMediaInfo:
				isMedia = true
				if media, e := newMedia(value); e != nil {
					err = nil
				} else {
					m.Media = append(m.Media, media)
				}
			default:
				m.UnsupportLines = append(m.UnsupportLines, content)
			}
		}
	}
}
