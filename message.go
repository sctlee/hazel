package tcpx

// import (
// 	"fmt"
// 	"strings"
// )
//
// type IMessage interface {
// 	Get() map[string]string // return map[string]string
// 	SetBoardClients(boardCids []ClientID)
// 	exec()
// 	send()
// 	bcast()
// }
//
// type Message struct {
// 	rawData     string
// 	client      ClientID
// 	multiClient []ClientID
// }
//
// func NewMessage(cid ClientID, data interface{}) IMessage {
// 	var d string
// 	switch data.(type) {
// 	case string:
// 		d = data.(string)
// 	case map[string]string:
// 		d = pt.UnMarshal(data.(map[string]string))
// 	default:
// 		d = "Error: Can't parse the type of message!"
// 	}
// 	return &Message{
// 		rawData: d,
// 		client:  cid,
// 	}
// }
//
// func NewBoardMessage(cid ClientID, data interface{}, boardCids []ClientID) IMessage {
// 	msg := NewMessage(cid, data)
// 	msg.SetBoardClients(boardCids)
// 	fmt.Println(msg)
// 	return msg
// }
//
// func (self *Message) Get() map[string]string {
// 	return pt.Marshal(self.rawData)
// }
//
// func (self *Message) SetBoardClients(boardCids []ClientID) {
// 	self.multiClient = make([]ClientID, 0)
// 	for _, cid := range boardCids {
// 		self.multiClient = append(self.multiClient, cid)
// 	}
// }
//
// func (self *Message) exec() {
// 	if self.multiClient != nil || len(self.multiClient) != 0 {
// 		self.bcast()
// 	} else {
// 		self.send()
// 	}
// }
// func (self *Message) send() {
// 	sname := strings.Split(string(self.client), ".")[0]
// 	if sname == serverName {
// 		GetClientByID(self.client).PutOutgoing(self.rawData)
// 	} else {
// 		PutMessage2Harbor(self.client, self.rawData)
// 	}
// }
// func (self *Message) bcast() {
// 	for _, client := range self.multiClient {
// 		go func(cid ClientID) {
// 			sname := strings.Split(string(cid), ".")[0]
// 			if sname == serverName {
// 				GetClientByID(cid).PutOutgoing(self.rawData)
// 			} else {
// 				PutMessage2Harbor(cid, self.rawData)
// 			}
// 		}(client)
// 	}
// }
//
// type MessageManager struct {
// 	receiver chan IMessage
// }
//
// var mmanager *MessageManager = NewMessageManager()
//
// func NewMessageManager() *MessageManager {
// 	mm := &MessageManager{
// 		receiver: make(chan IMessage, 10),
// 	}
//
// 	go func() {
// 		for m := range mm.receiver {
// 			go m.exec()
// 		}
// 	}()
//
// 	return mm
// }
//
// func SendMessage(message IMessage) {
// 	mmanager.receiver <- message
// }
