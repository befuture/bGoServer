package main

import (
	"agent"
	"fmt"
	"time"
)

func StaticControlPanel() {
	fmt.Println("Control panel")
	fmt.Printf("1. Show connection count \n2. Exit server\n")

	for {
		time.Sleep(10 * time.Second)
	}
}

func main() {
	agent.ServerStart()
	StaticControlPanel()

}
