package resp

// Reply 接口表示 RESP 协议的通用返回类型，每种实现需要提供 ToBytes 方法。
type Reply interface {
	ToBytes() []byte
}
