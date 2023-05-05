/*
问题:
在使用tcp通信的消息中， 消息头部由uint32的消息体长度、uint16的消息ID、uint16的服务ID、int8的校验值四个字段组成；消息体为protobuf序列化的数据， 使用gnet库搭建服务器，给出具体实现代码。
*/

package main

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/panjf2000/gnet"
)

type Message struct {
	BodyLength uint32 // 消息体长度
	MsgID      uint16 // 消息ID
	ServiceID  uint16 // 服务ID
	Checksum   int8   // 校验值
	Body       []byte // 消息体
}

type Server struct {
	*gnet.EventServer
}

func (es *Server) OnInitComplete(srv gnet.Server) (action gnet.Action) {
	fmt.Println("Server started ...")
	return
}

func (es *Server) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	// 检查数据是否满足消息头部长度
	if len(frame) < 10 {
		return
	}
	// 解析消息头部
	var message Message
	buffer := bytes.NewBuffer(frame)
	binary.Read(buffer, binary.BigEndian, &message.BodyLength)
	binary.Read(buffer, binary.BigEndian, &message.MsgID)
	binary.Read(buffer, binary.BigEndian, &message.ServiceID)
	binary.Read(buffer, binary.BigEndian, &message.Checksum)

	// 检查消息体长度是否正确
	if len(frame)-10 < int(message.BodyLength) {
		return
	}

	// 解析消息体
	message.Body = make([]byte, message.BodyLength)
	copy(message.Body, frame[10:10+message.BodyLength])
	var messageData proto.Message
	proto.Unmarshal(message.Body, messageData)

	// 处理消息体
	// ...

	return
}

func main() {
	err := gnet.Serve(&Server{}, "tcp://:9000", gnet.WithMulticore(true), gnet.WithCodec(gnet.NewLengthFieldBasedFrameCodec(binary.BigEndian, 0, 4, 0, 0)))
	if err != nil {
		panic(err)
	}
}
