package tcpx

import "log"

const (
	CLIENT_STATE_OPEN  = 1
	CLIENT_STATE_CLOSE = 2
)

type Xtime struct {
	isExist  bool
	question string
}

type IClient interface {
	TRead(incoming chan string) error
	TWrite(outgoing chan string) error
	Close()
}

type Client struct {
	c        IClient
	incoming chan string
	outgoing chan string
	State    int
}

func CreateClient(ic IClient) (client *Client) {
	client = &Client{
		c: ic,
		// Conn:     conn,
		incoming: make(chan string),
		outgoing: make(chan string),
		State:    CLIENT_STATE_OPEN,
	}

	go client.Read()
	go client.Write()

	return
}

func (self *Client) GetIncoming() chan string {
	return self.incoming
}

func (self *Client) PutOutgoing(str string) {
	if self.State == CLIENT_STATE_OPEN {
		self.outgoing <- str
	}
}

func (self *Client) Close() {
	self.c.Close()
	self.State = CLIENT_STATE_CLOSE
	// close mean to notify a receiver not to expect any more values to be sent.
	// but in a feature, it doesn't know the conn's stat, so it doesn't know if
	// the channel is useless, so it can't close the channel, so don't close it
	// here(it's not a producer)
	close(self.incoming)
	close(self.outgoing)
	logger.Println("Client close")
	log.Println("Client close")
}

func (self *Client) Read() {
	err := self.c.TRead(self.incoming)
	if err != nil {
		logger.Printf("Read error %s\n", err)
		log.Printf("Read error %s\n", err)
		self.Close()
	}
}

func (self *Client) Write() {
	err := self.c.TWrite(self.outgoing)
	if err != nil {
		logger.Printf("Write error %s\n", err)
		log.Printf("Write error %s\n", err)
	} else {
		log.Println("client writer closed")
		logger.Println("client writer closed")
	}
}

// func (client *Client) TRead() {
// 	reader := bufio.NewReader(client.Conn)
// 	for {
// 		if line, _, err := reader.ReadLine(); err == nil {
// 			client.Incoming <- string(line)
// 		} else {
// 			fmt.Printf("Read error: %s\n", err)
// 			return
// 		}
//
// 	}
// }
//
// func (client *Client) TWrite() {
// 	writer := bufio.NewWriter(client.Conn)
// 	for data := range client.Outgoing {
// 		writer.WriteString(data + "\n")
// 		// q: why flush is necessary? a:using buf mean: it won't send immedicately until buf is full
// 		writer.Flush()
// 	}
// }
