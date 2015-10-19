package tcpx

type Xtime struct {
	isExist  bool
	question string
}

type IClient interface {
	TRead(incoming chan string)
	TWrite(outgoing chan string)
	Close()
}

type Client struct {
	c        IClient
	incoming chan string
	outgoing chan string
}

func CreateClient(ic IClient) (client *Client) {
	client = &Client{
		c: ic,
		// Conn:     conn,
		incoming: make(chan string),
		outgoing: make(chan string),
	}

	go client.c.TRead(client.incoming)
	go client.c.TWrite(client.outgoing)

	return
}

func (self *Client) GetIncoming() chan string {
	return self.incoming
}

func (self *Client) PutOutgoing(str string) {
	self.outgoing <- str
}

func (self *Client) Close() {
	logger.Println("client close")
	self.c.Close()
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
