package tufer

import (
	"net"
)

type Connection struct {
	controlChannel *net.TCPConn
}

func ( c *Connection ) Close() {
	c.controlChannel.Close()
}
