package reply

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBulkReply(t *testing.T) {
	// 测试非空 BulkReply
	bulk := MakeBulkReply([]byte("hello"))
	expected := "$5\r\nhello\r\n"
	assert.Equal(t, expected, string(bulk.ToBytes()))

	// 测试空 BulkReply
	emptyBulk := MakeBulkReply(nil)
	expectedEmpty := "$-1\r\n"
	assert.Equal(t, expectedEmpty, string(emptyBulk.ToBytes()))
}

func TestMultiBulkReply(t *testing.T) {
	// 测试非空 MultiBulkReply
	multiBulk := MakeMultiBulkReply([][]byte{
		[]byte("foo"),
		[]byte("bar"),
		nil,
	})
	expected := "*3\r\n$3\r\nfoo\r\n$3\r\nbar\r\n$-1\r\n"
	assert.Equal(t, expected, string(multiBulk.ToBytes()))

	// 测试空 MultiBulkReply
	emptyMultiBulk := MakeMultiBulkReply(nil)
	expectedEmpty := "*0\r\n"
	assert.Equal(t, expectedEmpty, string(emptyMultiBulk.ToBytes()))
}

func TestStatusReply(t *testing.T) {
	// 测试状态回复
	status := MakeStatusReply("OK")
	expected := "+OK\r\n"
	assert.Equal(t, expected, string(status.ToBytes()))
}

func TestIntReply(t *testing.T) {
	// 测试整数回复
	intReply := MakeIntReply(123)
	expected := ":123\r\n"
	assert.Equal(t, expected, string(intReply.ToBytes()))
}

func TestStandardErrReply(t *testing.T) {
	// 测试标准错误回复
	errReply := MakeErrReply("Custom error")
	expected := "-Custom error\r\n"
	assert.Equal(t, expected, string(errReply.ToBytes()))
	assert.Equal(t, "Custom error", errReply.Error())
}

func TestIsErrorReply(t *testing.T) {
	// 测试错误回复判断
	errReply := MakeErrReply("Error")
	statusReply := MakeStatusReply("OK")
	assert.True(t, IsErrorReply(errReply))
	assert.False(t, IsErrorReply(statusReply))
}
