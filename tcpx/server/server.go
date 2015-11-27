package server

import (
	"fmt"
	"net"

	"github.com/sctlee/tcpx/daemon"
	"github.com/sctlee/tcpx/daemon/message"
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
	s      base.IServer
	d      *daemon.Daemon
	config *ServerConfig
	// listener net.Listener
	// quiting  chan net.Conn
	start    chan struct{}
	incoming chan string
	outgoing chan string

	joinedNumber int
	// Routers RouterList
}

func (self *Server) Start() error {
	self.s.Listen(self.config.Port)
	defer self.Close()

	<-self.start
	fmt.Println("server start")

	for {
		conn, err := self.s.Accept()
		if err != nil {
			fmt.Println(err)
			return err
		} else {

			genClientID := func() string {
				self.joinedNumber++
				return fmt.Sprintf("%s.%d", self.config.ServerName, self.joinedNumber)
			}

			cid := genClientID()

			fmt.Println("client id :" + cid)
			tempClient := client.CreateClient(conn, cid)
			self.d.Pending <- tempClient

			go func(c *client.Client) {
				defer func() {
					self.d.Quiting <- c
				}()

				for {
					rawData, ok := c.GetIncoming()
					if !ok {
						break
					}

					fmt.Println(rawData)

					err := self.d.MsgManager.PutMessage(
						message.NewMessage(
							self.config.Pt, string(cid), "", rawData, daemon.MESSAGE_TYPE_TOSERVICE))

					if err != nil {
						fmt.Sprintf("server.error|msg:%s", err)
						self.d.MsgManager.PutMessage(
							daemon.NewSimpleMessage(
								string(cid),
								fmt.Sprintf("server.error|msg:%s", err)))
					}
				}
			}(tempClient)
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

func NewServer(cf *ServerConfig) *Server {
	return &Server{
		s:        base.NewTCPServer(),
		config:   cf,
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
