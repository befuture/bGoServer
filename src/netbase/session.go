package netbase

import (
	"log"
	"net"
)

type Session struct {
	id          int
	HashID      string
	Conn        net.Conn
	IsAvaliable bool
	Nm          *Manager
	RecvBuffer  Buffer
	SendBuffer  Buffer
	GroupID     int
}

func (self *Session) GetID() int {
	return self.id
}

func (self *Session) Open(conn net.Conn, nm *Manager, id int) {
	self.IsAvaliable = true
	self.RecvBuffer.Init()
	self.SendBuffer.Init()
	self.Conn = conn
	self.Nm = nm
	self.id = id
}

func (self *Session) Close() {
	self.Nm.RemoveFromGroup(self.GroupID, self.GetID())
	self.Nm.ConnectionCloseChan <- -1
	self.IsAvaliable = false
	self.Conn.Close()
}

func (self *Session) SessionHandler() {
	//errFlag := false
	readBuffer := make([]byte, 1024)

	for {
		count, err := self.Conn.Read(readBuffer)
		if err != nil {
			log.Println("Read error" + err.Error())
			break
		}

		self.RecvBuffer.PushData(readBuffer[0:count])
		log.Printf("Read data length(%d), buffer length(%x)\n", len(readBuffer), len(self.RecvBuffer.Buffer.Bytes()))
		if !self.Nm.Parse(self) {
			log.Println("What's the fuck?!")
			self.Close()
			return
		}
	}
	defer self.Close()
}

func (self *Session) Send(sendBuffer []byte) {
	totalCount := len(sendBuffer)
	var totalSend int = 0

	for {
		count, err := self.Conn.Write(sendBuffer[totalSend:totalCount])
		totalSend += count
		switch {
		case err != nil:
			//! add a panic and add a safe hook with this...
			return
		case int64(totalSend) == int64(totalCount):
			return
		default:
			totalSend += count
		}
	}

	return
}
