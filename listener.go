package tufer


import (
	"net"
)


type Listener struct {
	listener *net.TCPListener
	tcpAddr  *net.TCPAddr
}


func NewListener(service string)( *Listener, error ) {

	tcpAddr, err := net.ResolveTCPAddr( "ip4", service )

	if err != nil {
		return nil, err
	}	

	listener, err := net.ListenTCP( "tcp", tcpAddr )

	if err != nil {
		return nil, err
	}

	return &Listener{listener, tcpAddr}, nil	

}

func ( l *Listener ) Accept()( *Connection, error ) {
	controlChannel, err := l.listener.Accept()
	
	if err != nil {
		return nil, err
	}

	return &Connection{controlChannel.(*net.TCPConn)}, nil
	 
}
