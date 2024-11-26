package reply

// PongReply 表示 RESP 的 "+PONG" 简单字符串回复。
type PongReply struct {
}

// 预定义的 "+PONG" 响应字节数组，避免重复分配内存。
var pongBytes = []byte("+PONG\r\n")

// ToBytes 将 PongReply 序列化为 RESP 格式的字节数组。
func (r PongReply) ToBytes() []byte {
	return pongBytes
}

// MakePongReply 返回一个 PongReply 实例。
func MakePongReply() *PongReply {
	return &PongReply{}
}

// OkReply 表示 RESP 的 "+OK" 简单字符串回复。
type OkReply struct {
}

// 预定义的 "+OK" 响应字节数组。
var okBytes = []byte("+OK\r\n")

// ToBytes 将 OkReply 序列化为 RESP 格式的字节数组。
func (r *OkReply) ToBytes() []byte {
	return okBytes
}

// 单例的 OkReply 实例，避免重复分配。
var theOkReply = new(OkReply)

// MakeOkReply 返回一个单例的 OkReply 实例。
func MakeOkReply() *OkReply {
	return theOkReply
}

// NullBulkReply 表示 RESP 的空字符串 "$-1" 回复。
type NullBulkReply struct {
}

// 预定义的 "$-1" 响应字节数组。
var nullBulkBytes = []byte("$-1\r\n")

// ToBytes 将 NullBulkReply 序列化为 RESP 格式的字节数组。
func (n NullBulkReply) ToBytes() []byte {
	return nullBulkBytes
}

// MakeNullBulkReply 返回一个 NullBulkReply 实例。
func MakeNullBulkReply() *NullBulkReply {
	return &NullBulkReply{}
}

// EmptyMultiBulkReply 表示 RESP 的空多条回复 "*0"。
type EmptyMultiBulkReply struct{}

// 预定义的 "*0" 响应字节数组。
var emptyMultiBulkBytes = []byte("*0\r\n")

// ToBytes 将 EmptyMultiBulkReply 序列化为 RESP 格式的字节数组。
func (r *EmptyMultiBulkReply) ToBytes() []byte {
	return emptyMultiBulkBytes
}

// NoReply 表示 RESP 的无返回情况，例如订阅命令。
type NoReply struct{}

// 预定义的空响应字节数组。
var noBytes = []byte("")

// ToBytes 将 NoReply 序列化为 RESP 格式的字节数组。
func (r *NoReply) ToBytes() []byte {
	return noBytes
}
