package sdp

// 媒体项
type Media struct {
	Info          MediaInfo   // m 媒体名称和传输地址
	Title         string      // i* 媒体标题
	Connection    *Connection // c* 连接信息，如果会话层有，这里可以没有
	UnsupportLine []string    // TODO 不支持处理的行
}

func newMedia(str string) (item Media, err error) {
	info, err := parseMediaInfo(str)
	if err != nil {
		return
	}
	item = Media{
		Info: info,
	}
	return
}

// 字符串表示
func (m Media) String() (result string) {
	result += output(FieldMediaInfo, m.Info.String())
	if len(m.Title) > 0 {
		result += output(FieldMediaTitle, m.Title)
	}
	if m.Connection != nil {
		result += output(FieldMediaConnection, m.Connection.String())
	}
	for _, line := range m.UnsupportLine {
		result += line + CRLF
	}
	return
}

// TODO delete 添加不支持的行
func (m *Media) addUnsupportLine(line string) {
	m.UnsupportLine = append(m.UnsupportLine, line)
}
