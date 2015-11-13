package tcpx

import (
	"fmt"
	"net"

	"github.com/sctlee/tcpx/base"
)

const (
	MAXCLIENTS = 50
)

type ClientTable map[IClient]*Client

type IServer interface {
	Listen(port string)
	Accept() (net.Conn, error)
	Close()
}

type Server struct {
	s base.IServer
	// listener net.Listener
	clients ClientTable
	pending chan *Client
	// quiting  chan net.Conn
	incoming chan string
	outgoing chan string

	Routers RouterList
}

func CreateServer() (server *Server) {
	server = &Server{
		s:       base.NewTCPServer(),
		clients: make(ClientTable),
		Routers: make(RouterList),
		pending: make(chan *Client),
		// quiting:  make(chan net.Conn),
		incoming: make(chan string),
		outgoing: make(chan string),
	}
	return
}

func (self *Server) Listen(port string) {
	go func() {
		for {
			select {
			case msg := <-self.incoming:
				fmt.Println(msg)
			case client := <-self.pending:
				self.Join(client)
				//
				// case conn := <-server.quiting:
			}
		}
	}()
	self.s.Listen(port)
}

func (self *Server) Join(client *Client) {
	self.clients[client.c] = client

	logger.Println("one client joined ")

	go func(c *Client) {
		defer func() {
			delete(self.clients, c.c)
			logger.Println("one client quited")
		}()

		for {
			msg, ok := c.GetIncoming()
			if !ok {
				break
			}
			if !self.Routers.RouteMsg(c, msg) {
				c.PutOutgoing("command error, Usage:'chatroom join 1','chatroom send hello'")
				// self.incoming <- msg
			}
		}
	}(client)
}

func (self *Server) Start(port string) {
	self.Listen(port)
	logger.Println("server start")
	// l, _ := net.Listen("tcp", fmt.Sprintf(":%s", port))
	// self.listener = l
	// defer self.listener.Close()
	defer self.Close()
	// chan listen

	for {
		conn, err := self.s.Accept()
		if err != nil {
			logger.Println(err)
		} else {
			client := CreateClient(conn)
			self.pending <- client
		}
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
	}
}

func (self *Server) Close() {
	logger.Println("server close")
	self.s.Close()
}
