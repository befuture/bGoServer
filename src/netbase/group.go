package netbase

import (
	"log"
)

type Group struct {
	id          int
	sessions    [1000]int
	isAvaliable bool
	count       int
}

func (self *Group) Init(id int) {
	self.id = 0
	self.isAvaliable = false
	self.count = 0
	for id, _ := range self.sessions {
		self.sessions[id] = -1
	}
}

func (self *Group) Join(session *Session) {
	for id, _ := range self.sessions {
		if self.sessions[id] == -1 {
			log.Printf("Joined @Stock [%d]\n", id)
			self.sessions[id] = session.GetID()
			self.count++
			break
		}
	}
}

func (self *Group) Remove(session *Session) {
	for id, _ := range self.sessions {
		if self.sessions[id] == session.GetID() {
			self.sessions[id] = -1
			self.count--
			break
		}
	}
}

func (self *Group) GetCount() int {
	return self.count
}

func (self *Group) IsEmpty() bool {
	return self.count == 0
}

func (self *Group) SetEmpty() {
	for id, _ := range self.sessions {
		self.sessions[id] = -1
	}
	self.count = 0
}
