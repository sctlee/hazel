package tcpx

import (
	"bufio"
	"fmt"
	"net"
)

type IClient interface {
	TRead(incoming chan string) error
	TWrite(outgoing chan string) error
	Close()
}

type TCPServer struct {
	listener net.Listener
}

type TCPClient struct {
	Conn net.Conn
}

func (self *TCPServer) Listen(port string) {
	self.listener, _ = net.Listen("tcp", fmt.Sprintf(":%s", port))
}

func (self *TCPServer) Accept() (tc *TCPClient) {
	if conn, err := self.listener.Accept(); err == nil {
		tc = &TCPClient{
			Conn: conn,
		}
	} else {
		fmt.Println(err)
	}
	return
}

func (self *TCPServer) Close() {
	self.listener.Close()
}

func (self *TCPClient) TRead(incoming chan string) error {
	reader := bufio.NewReader(self.Conn)
	for {
		if line, _, err := reader.ReadLine(); err == nil {
			incoming <- string(line)
		} else {
			// fmt.Printf("Read error: %s\n", err)
			self.Conn.Close()
			return err
		}
	}
}

func (self *TCPClient) TWrite(outgoing chan string) error {
	var err error
	writer := bufio.NewWriter(self.Conn)
	for data := range outgoing {
		_, err = writer.WriteString(data + "\n")
		// q: why flush is necessary? a:using buf mean: it won't send immedicately until buf is full
		if err != nil {
			break
		}
		writer.Flush()
	}
	return err
}

func (self *TCPClient) Close() {
	self.Conn.Close()
}
