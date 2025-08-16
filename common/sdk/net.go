package sdk

import (
	"encoding/json"
	"fmt"
	"net"
	"sync/atomic"

	"github.com/gogo/protobuf/proto"
	"github.com/haozhe-qi/ChatIm/common/idl/message"
	"github.com/haozhe-qi/ChatIm/common/tcp"
)

type connect struct {
	sendChan, recvChan chan *Message
	conn               *net.TCPConn
	connID             uint64
	ip                 net.IP
	port               int
}

func newConnet(ip net.IP, port int) *connect {
	clientConn := &connect{
		sendChan: make(chan *Message),
		recvChan: make(chan *Message),
		ip:       ip,
		port:     port,
	}
	addr := &net.TCPAddr{IP: ip, Port: port}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		fmt.Printf("DialTCP.err=%+v", err)
		return nil
	}
	clientConn.conn = conn
	return clientConn
}
func handAckMsg(c *connect, data []byte) *Message {
	ackMsg := &message.ACKMsg{}
	proto.Unmarshal(data, ackMsg)
	switch ackMsg.Type {
	case message.CmdType_Login, message.CmdType_ReConn:
		atomic.StoreUint64(&c.connID, ackMsg.ConnID)
	}
	return &Message{
		Type:       MsgTypeAck,
		Name:       "plato",
		FromUserID: "1212121",
		ToUserID:   "222212122",
		Content:    ackMsg.Msg,
	}
}
func handPushMsg(c *connect, data []byte) *Message {
	pushMsg := &message.PushMsg{}
	proto.Unmarshal(data, pushMsg)
	// if pushMsg.MsgID == c.maxMsgID+1 {
	// 	c.maxMsgID++
	msg := &Message{}
	json.Unmarshal(pushMsg.Content, msg)
	ackMsg := &message.ACKMsg{
		Type:   message.CmdType_UP,
		ConnID: c.connID,
	}
	ackData, _ := proto.Marshal(ackMsg)
	c.send(message.CmdType_ACK, ackData)
	return msg

}
func (c *connect) reConn() {
	c.conn.Close()
	addr := &net.TCPAddr{IP: c.ip, Port: c.port}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		fmt.Printf("DialTCP.err=%+v", err)
	}
	c.conn = conn
}
func (c *connect) send(ty message.CmdType, palyload []byte) {
	// 直接发送给接收方
	msgCmd := message.MsgCmd{
		Type:    ty,
		Payload: palyload,
	}
	msg, err := proto.Marshal(&msgCmd)
	if err != nil {
		panic(err)
	}
	dataPgk := tcp.DataPgk{
		Data: msg,
		Len:  uint32(len(msg)),
	}
	c.conn.Write(dataPgk.Marshal())
}
func (c *connect) recv() <-chan *Message {
	return c.recvChan
}

func (c *connect) close() {
	c.conn.Close()
}
