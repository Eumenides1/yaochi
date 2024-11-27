package reply

// ErrorReply 接口，表示 RESP 协议中的错误回复类型。
// 每个错误回复需要实现 ToBytes 方法用于序列化为 RESP 格式，
// 以及 Error 方法返回错误的字符串描述。
type ErrorReply interface {
	Error() string   // 返回错误信息的描述
	ToBytes() []byte // 将错误序列化为 RESP 格式字节数组
}
