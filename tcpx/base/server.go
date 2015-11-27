package base

import (
	"fmt"
	"net"
)

type TCPServer struct {
	listener net.Listener
}
type IServer interface {
	Listen(port string)
	Accept() (net.Conn, error)
	Close()
}

func NewTCPServer() IServer {
	return &TCPServer{}
}

func (self *TCPServer) Listen(port string) {
	self.listener, _ = net.Listen("tcp", fmt.Sprintf(":%s", port))
}

func (self *TCPServer) Accept() (net.Conn, error) {
	conn, err := self.listener.Accept()
	return conn, err
}

func (self *TCPServer) Close() {
	self.listener.Close()
}
