package sdp

import (
	"errors"
	"strings"
)

// 分隔符
const CRLF = "\r\n"
const seperator = " "

// 解析字段
func parse(str string) (key string, value string, err error) {
	loc := strings.Index(str, "=")
	if loc < 0 {
		err = errors.New("sdp: line format error")
		return
	}
	key = str[0:loc]
	value = str[loc+1:]
	return
}

// 单行输出
func output(key string, value string) string {
	return key + "=" + value + CRLF
}
