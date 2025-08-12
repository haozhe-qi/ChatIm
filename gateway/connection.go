package gateway

import "net"

type connection struct {
	fd   int //fd不能唯一标识会复用得修改
	conn *net.TCPConn
}

func (c *connection) Close() {
	err := c.conn.Close()
	panic(err)
}

func (c *connection) RemoteAddr() string {
	return c.conn.RemoteAddr().String()
}
