package tufer

import (
"net"

)


type Client struct {
	serverAddr *net.TCPAddr
}



func NewClient( service string  ) ( *Client, error ) {

	serverAddr, err := net.ResolveTCPAddr( "tcp4", service )
	
	if err != nil {
		return nil, err
	}

	return &Client{ serverAddr }, nil

}

func ( c *Client ) Connect() ( *Connection, error ) {

	controlChannel, err := net.DialTCP("tcp", nil, c.serverAddr )

	if err != nil {
		return nil, err
	}

	return &Connection{ controlChannel }, nil
}
