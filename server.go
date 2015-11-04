package tcpx

import (
	"fmt"
)

const (
	MAXCLIENTS = 50
)

type IServer interface {
	Listen(port string)
	Accept() (tc *TCPClient)
	Close()
}

type ClientTable map[IClient]*Client

type Server struct {
	s IServer
	// listener net.Listener
	clients ClientTable
	Routers RouterList
	pending chan IClient
	// quiting  chan net.Conn
	incoming chan string
	outgoing chan string
}

func CreateServer() (server *Server) {
	server = &Server{
		s:       &TCPServer{},
		clients: make(ClientTable),
		Routers: make(RouterList),
		pending: make(chan IClient),
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

func (self *Server) Join(ic IClient) {
	client := CreateClient(ic)
	self.clients[ic] = client

	logger.Println("one client joined ")

	go func(c *Client) {
		defer func() {
			delete(self.clients, c.c)
			logger.Println("one client quited")
		}()

		for {
			msg, ok := c.GetMessage()
			if !ok {
				break
			}
			// package msg whish conn
			// msg = fmt.Sprintf("format string", a ...interface{})
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
		self.pending <- self.s.Accept()
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
