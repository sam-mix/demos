package msg

// 客户端消息头
type ClientMsgHead struct {
	MsgLen uint32 // 消息体长度
	MsgID  uint16 // 消息ID
	SvrID  uint16 // 服务ID
	CCode  int8   // 校验值
}

const CLIMSGHEADLEN = 2 + 4 + 2 + 1
const MAXBODYLEN = 1024 * 8
