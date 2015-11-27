package base

import (
	"bufio"
	"net"
)

type TCPClient struct {
	Conn net.Conn
}

func NewTCPClient(conn net.Conn) *TCPClient {
	return &TCPClient{conn}
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
