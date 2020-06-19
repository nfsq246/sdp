# sdp

这是使用Go实现的RFC4566-SDP协议。多用于SIP、WebRTC等多媒体会话场景。

主要结构体为`Message`，对应一条完整的SDP信息。

### 使用方法

在Go中引用本库：

```go
import "github.com/nfsq246/sdp"
```

生成一个消息对象，解析外部收到的字符串：

```go
sdpMsg, err := sdp.NewMessage(transferString)
```

将消息对象转换为字符串用于传输：

```go
transferString := sdpMsg.String()
```

### TODO

- [ ] 增加完整的字段支持。

### change

增加了删除前后多余垃圾字符

### 参考资料

* [RFC4566-IETF](https://tools.ietf.org/html/rfc4566)
* [RFC4566-中文版](http://www.docin.com/p-52701644.html)
