package main

import (
	"demos/msg"
	"demos/pb"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"time"

	"google.golang.org/protobuf/proto"
)

func randString(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}

// 发送心跳请求
func sendHeartbeat(conn net.Conn) error {

	// 创建心跳请求消息
	c2s := pb.C2SLogin{
		Acc: randString(10),
	}

	// 将请求消息序列化为字节数组
	data, err := proto.Marshal(&c2s)
	if err != nil {
		return err
	}
	head := new(msg.ClientMsgHead)
	head.MsgLen = uint32(len(data))
	head.MsgID = uint16(pb.MSG_ID_LOGIN)
	head.SvrID = uint16(0)
	head.CCode = int8(0)
	headBuf, err := head.MsgPackBE()
	if err != nil {
		// log.Fatalln("proto.Marshal error:", err)
		log.Println("proto.Marshal error:", err)
	}
	messageBuf := append(headBuf, data...)
	// 发送心跳请求消息
	_, err = conn.Write(messageBuf)
	if err != nil {
		return err
	}

	// fmt.Println("Sent a heartbeat request")

	return nil
}

// 接收心跳响应
func receiveHeartbeat(conn net.Conn) error {
	headBuf, err := RecvMessage(conn, msg.CLIMSGHEADLEN)
	if err != nil {
		// log.Fatalln(err)
		log.Println(err)
	}
	headMsg := new(msg.ClientMsgHead)
	err = headMsg.UnMsgPackBE(headBuf)
	if err != nil {
		// log.Fatalln(err)
		log.Println(err)
	}
	fmt.Println("headMsg:", headMsg)
	bodyBuf, err := RecvMessage(conn, int(headMsg.MsgLen))
	if err != nil {
		// log.Fatalln(err)
		log.Println(err)
	}

	// 反序列化响应消息
	s2c := pb.S2CLogin{}
	err = proto.Unmarshal(bodyBuf, &s2c)
	if err != nil {
		// panic(err)
		// log.Fatalln(err)
		log.Println(err)
		return nil
	}

	// 输出响应时间戳
	fmt.Printf("Received a heartbeat response at %#v\n", s2c)

	return nil
}

// 客户端连接
func connect() error {
	// 连接服务器
	conn, err := net.Dial("tcp", "192.168.0.219:9999")
	if err != nil {
		return err
	}

	// 发送/接收心跳消息
	for {
		// go sendHeartbeat(conn)
		// go receiveHeartbeat(conn)
		sendHeartbeat(conn)
		receiveHeartbeat(conn)
		time.Sleep(time.Second * 5)
	}
}

func RecvMessage(conn net.Conn, nLen int) ([]byte, error) {
	defer func() {
		if recoverErr := recover(); recoverErr != nil {
			log.Println(recoverErr)
		}
	}()

	bodyBuff := make([]byte, nLen)
	bodyNum, err := io.ReadFull(conn, bodyBuff)
	if err != nil {
		return nil, err
	}

	if bodyNum != nLen {
		return nil, errors.New("recv data error: lesss than len ")
	}
	return bodyBuff, nil
}

// 主函数
func main() {
	err := connect()
	if err != nil {
		log.Println(err)
		// panic(err)
	}
}
