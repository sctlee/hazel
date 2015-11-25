package tcpx

import (
	"fmt"
	"net"

	"github.com/sctlee/tcpx/base"
)

const (
	MAXCLIENTS = 50
)

type ClientID string
type ClientTable map[ClientID]*Client

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
	incoming     chan string
	outgoing     chan string
	joinedNumber int

	Routers RouterList
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
	genClientID := func() ClientID {
		return ClientID(fmt.Sprintf("%s.%d", serverName, self.joinedNumber))
	}

	cid := genClientID()
	fmt.Println("client id :" + cid)
	self.clients[cid] = client
	self.joinedNumber++
	logger.Println("one client joined ")

	go func(cid ClientID) {
		defer func() {
			delete(self.clients, cid)
			logger.Println("one client quited")
		}()
		c := self.clients[cid]

		for {
			msg, ok := c.GetIncoming()
			if !ok {
				break
			}
			if !self.Routers.RouteMsg(cid, msg) {
				c.PutOutgoing("command error, Usage:'chatroom join 1','chatroom send hello'")
				// self.incoming <- msg
			}
		}
	}(cid)
}

func (self *Server) Start(port string) {
	self.Listen(port)
	logger.Println("server start")
	defer self.Close()

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

// func (self *Server) RegisterService(service *Services) {
//
// }

func (self *Server) Close() {
	logger.Println("server close")
	self.s.Close()
}

var server *Server

func CreateMainServer() *Server {
	server = CreateServer()

	// start harbor
	HarborStart()

	return server
}
func CreateServer() *Server {
	return &Server{
		s:       base.NewTCPServer(),
		clients: make(ClientTable),
		Routers: make(RouterList),
		pending: make(chan *Client),
		// quiting:  make(chan net.Conn),
		incoming: make(chan string),
		outgoing: make(chan string),
	}
}

func GetClientByID(cid ClientID) *Client {
	return server.clients[cid]
}
