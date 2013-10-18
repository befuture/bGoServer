package netbase

type IPacket interface {
	Parse(nm *Manager, ns *Session) (length int, err error)
	GetLength(buff []byte) (length int, err error)
}

type IPacketCutter interface {
	CutPacket(nm *Manager, buff []byte) (packetHandler *IPacket, err bool)
}

/*
type CommonPacket struct {
	Length int
}

func NewCommonPacket() *CommandPacket {
	packet := new(CommonPacket)
	packet.Length = -1;
	return packet
}

func (self *CommonPacket) Parse(nm *Manager, ns *Session) (count int, err error) {
	packetlen, _ := ns.RecvBuffer.ReadWord()
	packetid, _ := ns.RecvBuffer.ReadWord()
	packetdata, _ := ns.RecvBuffer.ReadBytes(self.Length)

	return self.Length, nil
}

func (self *CommandPacket) GetLength(buff []byte) (count int, err error) {
	return self.Length + 3, nil
}
*/
