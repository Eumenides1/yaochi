package reply

// UnknownErrReply 表示未知错误的回复类型
type UnknownErrReply struct{}

// 预定义的 RESP 协议 "+Err unknown" 的字节数组
var unknownErrBytes = []byte("-Err unknown\r\n")

// ToBytes 将 UnknownErrReply 序列化为 RESP 格式
func (r *UnknownErrReply) ToBytes() []byte {
	return unknownErrBytes
}

// Error 返回错误的字符串描述
func (r *UnknownErrReply) Error() string {
	return "Err unknown"
}

// ArgNumErrReply 表示命令参数数量错误的回复类型
type ArgNumErrReply struct {
	Cmd string // 存储命令名称
}

// ToBytes 将 ArgNumErrReply 序列化为 RESP 格式
func (r *ArgNumErrReply) ToBytes() []byte {
	return []byte("-ERR wrong number of arguments for '" + r.Cmd + "' command\r\n")
}

// Error 返回错误的字符串描述
func (r *ArgNumErrReply) Error() string {
	return "ERR wrong number of arguments for '" + r.Cmd + "' command"
}

// MakeArgNumErrReply 工厂方法，创建 ArgNumErrReply 实例
func MakeArgNumErrReply(cmd string) *ArgNumErrReply {
	return &ArgNumErrReply{
		Cmd: cmd,
	}
}

// SyntaxErrReply 表示语法错误的回复类型
type SyntaxErrReply struct{}

// 预定义的 RESP 协议 "-Err syntax error" 的字节数组
var syntaxErrBytes = []byte("-Err syntax error\r\n")

// 单例模式的语法错误回复实例，避免重复分配
var theSyntaxErrReply = &SyntaxErrReply{}

// MakeSyntaxErrReply 工厂方法，返回单例的 SyntaxErrReply 实例
func MakeSyntaxErrReply() *SyntaxErrReply {
	return theSyntaxErrReply
}

// ToBytes 将 SyntaxErrReply 序列化为 RESP 格式
func (r *SyntaxErrReply) ToBytes() []byte {
	return syntaxErrBytes
}

// Error 返回错误的字符串描述
func (r *SyntaxErrReply) Error() string {
	return "Err syntax error"
}

// WrongTypeErrReply 表示操作类型错误的回复类型
type WrongTypeErrReply struct{}

// 预定义的 RESP 协议 "-WRONGTYPE" 的字节数组
var wrongTypeErrBytes = []byte("-WRONGTYPE Operation against a key holding the wrong kind of value\r\n")

// ToBytes 将 WrongTypeErrReply 序列化为 RESP 格式
func (r *WrongTypeErrReply) ToBytes() []byte {
	return wrongTypeErrBytes
}

// Error 返回错误的字符串描述
func (r *WrongTypeErrReply) Error() string {
	return "WRONGTYPE Operation against a key holding the wrong kind of value"
}

// ProtocolErrReply 表示协议错误的回复类型
type ProtocolErrReply struct {
	Msg string // 错误信息描述
}

// ToBytes 将 ProtocolErrReply 序列化为 RESP 格式
func (r *ProtocolErrReply) ToBytes() []byte {
	return []byte("-ERR Protocol error: '" + r.Msg + "'\r\n")
}

// Error 返回错误的字符串描述
func (r *ProtocolErrReply) Error() string {
	return "ERR Protocol error: '" + r.Msg + "'"
}
