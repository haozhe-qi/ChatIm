package sdk

type connect struct {
	serverAddr         string
	sendChan, recvChan chan *Message
}

func newConnect(serverAddr string) *connect {
	return &connect{
		serverAddr: serverAddr,
		sendChan:   make(chan *Message),
		recvChan:   make(chan *Message),
	}
}

func (c *connect) send(data *Message) {
	//直接发送给对方
	c.recvChan <- data
}
func (c *connect) recv() <-chan *Message {
	return c.recvChan
}

func (c *connect) close() {
	//目前没啥回收的
}
