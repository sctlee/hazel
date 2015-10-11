package tcpx

import (
	"bufio"
	"fmt"
	"net"
)

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
		fmt.Printf("%v", conn)
		tc = &TCPClient{
			Conn: conn,
		}
	} else {
		fmt.Println(err)
	}
	return
}

func (self *TCPServer) Close() {
	logger.Println("server close")
	self.listener.Close()
}

func (self *TCPClient) TRead(incoming chan string) {
	reader := bufio.NewReader(self.Conn)
	for {
		if line, _, err := reader.ReadLine(); err == nil {
			incoming <- string(line)
		} else {
			fmt.Printf("Read error: %s\n", err)
			return
		}
	}
}

func (self *TCPClient) TWrite(outgoing chan string) {
	writer := bufio.NewWriter(self.Conn)
	for data := range outgoing {
		writer.WriteString(data + "\n")
		// q: why flush is necessary? a:using buf mean: it won't send immedicately until buf is full
		writer.Flush()
	}
}
