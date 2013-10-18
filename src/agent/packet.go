package agent

import (
	"log"
	"netbase"
)

type CommonPacket struct {
	Length int
}

func NewCommonPacket(len int) *CommonPacket {
	packet := new(CommonPacket)
	packet.Length = len
	return packet
}

func (self *CommonPacket) Parse(nm *netbase.Manager, ns *netbase.Session) (count int, err error) {
	packetid, _ := ns.RecvBuffer.ReadWord()
	packetseed, _ := ns.RecvBuffer.ReadByte()
	packetdata, _ := ns.RecvBuffer.ReadBytes(self.Length)
	log.Printf("Recv a packet (%x)\n", packetid)

	switch packetseed {
	case 'a':
		ns.Send(packetdata)
	case 'b':
		nm.SendToAll(packetdata)
	case 'c':
		nm.SendToAllExceptMe(packetdata, ns.GetID())
	case 'd':
		nm.SendToGroup(packetdata, ns.GroupID)
	case 'e':
		nm.SendToGroupExceptMe(packetdata, ns.GroupID, ns.GetID())
	}
	return self.Length + 3, nil
}

func (self *CommonPacket) GetLength(buff []byte) (count int, err error) {
	return self.Length + 3, nil
}

type DJMaxPacketCutter struct {
}

func (self *DJMaxPacketCutter) CutPacket(nm *netbase.Manager, buff []byte) (packetHandler *netbase.IPacket, err bool) {
	if len(buff) > 2 {
		var nPacketID int32 = int32(buff[0]) + int32(buff[1])*0x100
		if nm.PacketArray[nPacketID] != nil {
			length, err := (*nm.PacketArray[nPacketID]).GetLength(buff)
			//log.Printf("CutPacket(), Recv a packet (%x), length(%d)\n", nPacketID, length)
			if length <= len(buff) && err == nil {
				return nm.PacketArray[nPacketID], false
			}
		} else {
			//Unknow packet
			log.Printf("Got an unknow packet, disconnected! [%x]\n", nPacketID)
			//ns.Close()
			return nil, true
		}
	}

	return nil, false
}
