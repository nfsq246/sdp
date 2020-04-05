package sdp

import (
	"errors"
	"regexp"
	"strings"
)

// 媒体名称和传输地址
type MediaInfo struct {
	MediaType string // 媒体类型
	Port      string // 端口或端口列表
	Protocol  string // 传输协议
	Format    string // 其他格式，中间可以有空格
}

func parseMediaInfo(str string) (item MediaInfo, err error) {
	item = MediaInfo{}
	err = item.parse(str)
	return
}

// 字符串表达
func (mi MediaInfo) String() string {
	arr := []string{
		mi.MediaType,
		mi.Port,
		mi.Protocol,
		mi.Format,
	}
	return strings.Join(arr, seperator)
}

// 解析内容
func (mi *MediaInfo) parse(str string) (err error) {
	result := mediaInfoRegExp.FindStringSubmatch(str)
	if len(result) != 5 {
		err = errors.New("sdp: media info format error")
		return
	}
	mi.MediaType = result[1]
	mi.Port = result[2]
	mi.Protocol = result[3]
	mi.Format = result[4]
	return
}

var mediaInfoRegExp = regexp.MustCompile("^([^\\s]+) ([^\\s]+) ([^\\s]+) (.*)$")
