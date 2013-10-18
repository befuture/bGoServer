package agent

import (
	//"db"
	"log"
	"netbase"
	"strconv"
	//"time"
)

type ServerConfig struct {
	LocalPort     int
	MaxConnection int
	DbUsername    string
	DbPassword    string
	DbHostname    string
	DbDatabase    string
	MaxGroup      int
}

func (self *ServerConfig) LoadConfig() {
	log.Println("Loading server config file...")
	self.LocalPort = 50521
	self.MaxConnection = 1000
	self.DbUsername = "root"
	self.DbPassword = "ilovedjmax"
	self.DbHostname = "localhost"
	self.DbDatabase = "techBlog"
	self.MaxGroup = 1000
	log.Println("LocalPort: " + strconv.Itoa(self.LocalPort) + ", MaxConnection: " + strconv.Itoa(self.MaxConnection))
}

func ServerStart() {
	var sc ServerConfig
	var nm netbase.Manager

	log.Println("*******************************")
	log.Println("Beatstage Server v1.0")
	log.Println("Agent & Channel all in one mode")
	log.Println("*******************************")
	sc.LoadConfig()
	nm.SetPacketCutter(new(DJMaxPacketCutter))
	nm.RegistPacket(0x4141, NewCommonPacket(7))

	go nm.StartService(sc.LocalPort, sc.MaxConnection, sc.MaxGroup)
	/*
		for {
			time.Sleep(1 * time.Second)
		}
	*/
}
