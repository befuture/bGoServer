package netbase

import (
	"log"
	"net"
	"strconv"
)

type IManagerCallbackFunc interface {
	OnAccept(nm *Manager, ns *Session)
	OnConnected(nm *Manager, ns *Session)
	OnClose(nm *Manager, ns *Session)
}

type Manager struct {
	Sessions            []Session
	ConnectionCount     int
	ConnectionCloseChan chan int
	PacketArray         [65535]*IPacket
	PacketCutter        *IPacketCutter
	Groups              []Group
	Callback            IManagerCallbackFunc
}

func (self *Manager) Init(maxConnection int, maxGroup int) {
	self.ConnectionCount = 0
	self.Sessions = make([]Session, maxConnection)
	self.Groups = make([]Group, maxGroup)

	for key, _ := range self.Groups {
		self.Groups[key].Init(key)
	}
}

func (self *Manager) SendToAll(buff []byte) {
	for key, _ := range self.Sessions {
		if self.Sessions[key].IsAvaliable == true {
			go self.Sessions[key].Send(buff)
		}
	}
}

func (self *Manager) SendToGroup(buff []byte, groupid int) {
	for key, _ := range self.Groups[groupid].sessions {
		sessionId := self.Groups[groupid].sessions[key]
		if sessionId != -1 {
			self.Sessions[sessionId].Send(buff)
		}
	}
}

func (self *Manager) SendToGroupExceptMe(buff []byte, groupid, id int) {
	for key, _ := range self.Groups[groupid].sessions {
		sessionId := self.Groups[groupid].sessions[key]
		if sessionId != -1 && sessionId != id {
			self.Sessions[sessionId].Send(buff)
		}
	}
}

func (self *Manager) SendToAllExceptMe(buff []byte, id int) {
	for key, _ := range self.Sessions {
		if self.Sessions[key].IsAvaliable == true && key != id {
			go self.Sessions[key].Send(buff)
		}
	}
}

func (self *Manager) JoinGroup(groupid, sessionid int) {
	log.Printf("Join Group [%d], session id: %d\n", groupid, sessionid)
	self.Groups[groupid].Join(&self.Sessions[sessionid])
	self.Sessions[sessionid].GroupID = groupid
}

func (self *Manager) FindEmptyGroup() int {
	for key, _ := range self.Groups {
		if self.Groups[key].IsEmpty() {
			return key
		}
	}
	return -1
}

func (self *Manager) RemoveFromGroup(groupid, sessionid int) {
	self.Groups[groupid].Remove(&self.Sessions[sessionid])
	self.Sessions[sessionid].GroupID = 0
}

func (self *Manager) RegistPacket(PacketId int, packet IPacket) {
	self.PacketArray[PacketId] = &packet
}

func (self *Manager) SetPacketCutter(PacketCutter IPacketCutter) {
	self.PacketCutter = &PacketCutter
}

func (self *Manager) Parse(ns *Session) bool {
	for {
		buffer := ns.RecvBuffer.Buffer.Bytes()
		if len(buffer) == 0 {
			break
		}
		packetHandler, err := (*self.PacketCutter).CutPacket(self, buffer)
		if packetHandler == nil {
			if err {
				log.Println("Split packet error!")
				return false
			}
		} else {
			(*packetHandler).Parse(self, ns)
			log.Printf("PacketLength(%d), \n", len(ns.RecvBuffer.Buffer.Bytes()))
		}
	}
	return true
}

func (self *Manager) StaticConnection() {
	for countChange := range self.ConnectionCloseChan {
		self.ConnectionCount += countChange
		log.Println("Current connections: " + strconv.Itoa(self.ConnectionCount))
	}
}

func (self *Manager) StartService(port, maxconnection, maxgroup int) {

	self.Init(maxconnection, maxgroup)
	listener, err := net.Listen("tcp", "0.0.0.0:"+strconv.Itoa(port))

	if err != nil {
		log.Println("Error listening:" + err.Error())
		return
		//we be we should make some panic
	}

	log.Println("Service Start... [" + strconv.Itoa(port) + "], MaxConnect[" + strconv.Itoa(len(self.Sessions)) + ", " + strconv.Itoa(cap(self.Sessions)) + "]")

	self.ConnectionCloseChan = make(chan int)
	go self.StaticConnection()

	for {
		conn, err := listener.Accept()

		if err != nil {
			println("Error accept:", err.Error())
		}

		for key, _ := range self.Sessions {
			if self.Sessions[key].IsAvaliable == false {
				//Accept the connection
				log.Println("Open connection key: ", key)
				self.Sessions[key].Open(conn, self, key)
				self.JoinGroup(0, key)
				log.Println("Joined?")
				go self.Sessions[key].SessionHandler()
				self.ConnectionCloseChan <- 1
				break
			}
		}
	}

}
