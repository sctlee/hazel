package server

import (
	"fmt"
	"net"

	"github.com/sctlee/tcpx/daemon"
	"github.com/sctlee/tcpx/tcpx/base"
	"github.com/sctlee/tcpx/tcpx/client"
)

const (
	MAXCLIENTS = 50
)

type IServer interface {
	Listen(port string)
	Accept() (net.Conn, error)
	Close()
}

type Server struct {
	s base.IServer
	d *daemon.Daemon
	// listener net.Listener
	// quiting  chan net.Conn
	start    chan struct{}
	incoming chan string
	outgoing chan string

	// Routers RouterList
}

func (self *Server) Start(port string) error {
	self.s.Listen(port)
	defer self.Close()

	<-self.start
	fmt.Println("server start")

	for {
		conn, err := self.s.Accept()
		if err != nil {
			fmt.Println(err)
			return err
		} else {
			tempClient := client.CreateClient(conn)
			self.d.Pending <- tempClient
		}
	}
}

func (self *Server) AcceptConnections(daemon *daemon.Daemon) {
	self.d = daemon
	close(self.start)
}

func (self *Server) Close() {
	fmt.Println("server close")
	self.s.Close()
}

func NewServer() *Server {
	return &Server{
		s: base.NewTCPServer(),
		// quiting:  make(chan net.Conn),
		start:    make(chan struct{}),
		incoming: make(chan string),
		outgoing: make(chan string),
	}
}

// var server *Server

// func CreateMainServer() *Server {
// 	server = CreateServer()
//
// 	// start harbor
// 	HarborStart()
//
// 	return server
// }
//
// func GetClientByID(cid ClientID) *Client {
// 	return server.clients[cid]
// }

// if conn, err := self.listener.Accept(); err == nil {
// 	self.pending <- conn
// go func(c net.Conn) {
// 	buf := make([]byte, 1024)
// 	for {
// 		cn, err := c.Read(buf)
// 		if err != nil {
// 			c.Close()
// 			break
// 		}
// 		log.Println(cn, string(buf[:cn]))
// 	}
// }(conn)
// }
