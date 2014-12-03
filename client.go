package tufer

import (
"net"

)


type Client struct {
	serverAddr *net.TCPAddr
	localAddr  *net.TCPAddr
}



func NewClient( service string, laddr string  ) ( *Client, error ) {

	var localAddr *net.TCPAddr
	
	if len(laddr) > 0 {
		
		la, err := net.ResolveTCPAddr( "tcp4", laddr )
		
		if err != nil {
			return nil, err
		}

		localAddr = la
	}

	serverAddr, err := net.ResolveTCPAddr( "tcp4", service )
	
	if err != nil {
		return nil, err
	}

	return &Client{ serverAddr, localAddr }, nil

}

func ( c *Client ) Connect() ( *Connection, error ) {

	controlChannel, err := net.DialTCP("tcp", c.localAddr, c.serverAddr )

	if err != nil {
		return nil, err
	}

	return &Connection{ controlChannel }, nil
}
