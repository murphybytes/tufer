package tufer


import (
	"net"
)


type Listener struct {
	listener *net.TCPListener

}


func NewListener(service string)( *Listener, error ) {
	tcpAddr, err := net.ResolveTCPAddr( "ip4", service )

	if err != nil {
		return nil, err
	}
	
	l := new(Listener)
	listener, err := net.ListenTCP( "tcp", tcpAddr )
	l.listener = listener

	if err != nil {
		return nil, err
	}

	return l, nil	

}

func ( l *Listener ) Accept()( net.Conn, error ) {
	 return l.listener.Accept()
}
