package sdp

const (
	// 会话描述标识
	Description = "application/sdp"

	// 当前版本号
	Version = "0" // 默认版本，无小版本号

	// 地址类型
	AddressTypeIPv4 = "IP4" // IPv4
	AddressTypeIPv6 = "IP6" // IPv6

	// 网络类型常量
	NetworkTypeInternet = "IN" // 因特网

	// 传输协议
	ProtocolUdp     = "UDP"      // UDP
	ProtocolRtpAvp  = "RTP/AVP"  // RTP/AVP
	ProtocolRtpSavp = "RTP/SAVP" // RTP/SAVP

	// 媒体类型
	MediaTypeAudio       = "audio"       // 音频
	MediaTypeVideo       = "video"       // 视频
	MediaTypeText        = "text"        // 文本
	MediaTypeApplication = "application" // 应用
	MediaTypeMessage     = "message"     // 消息

	// 会话层，描述字段按顺序排列
	FieldVersion     = "v" // 协议版本
	FieldBasic       = "o" // 发起者和会话标识符
	FieldSessionName = "s" // 会话名字
	FieldSessionInfo = "i" // * 会话信息
	FieldURI         = "u" // * URI标识符
	FieldEmail       = "e" // * 邮件地址
	FieldPhone       = "p" // * 电话号码
	FieldConnection  = "c" // * 连接信息，只有所有媒体层都有这个字段时，会话层才可以没有
	FieldBandwidth   = "b" // * 0个或多个带宽信息
	// 一个或多个时间描述
	FieldTimeZone      = "z" // * 时域调整
	FieldEncryptionKey = "k" // * 加密密钥
	FieldAttribute     = "a" // * 0个或多个会话属性
	// 0个或多个媒体描述

	// 时间描述字段
	FieldTime        = "t" // 会话有效的时间
	FieldRepeatTimes = "r" // * 0个或多个重复时间

	// 媒体描述字段，如果存在的话
	FieldMediaInfo          = "m" // 媒体名字和传输地址
	FieldMediaTitle         = "i" // * 媒体标题
	FieldMediaConnection    = "c" // * 连接信息，如果在会话层出现，则这里可以不出现
	FieldMediaBandwidth     = "b" // * 0个或多个带宽信息
	FieldMediaEncryptionKey = "k" // * 加密密钥
	FieldMediaAttribute     = "a" // * 0个或多个媒体属性
)
