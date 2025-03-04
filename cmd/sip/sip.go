package main

import (
	"fmt"

	pcksip "github.com/totqwq/medium/pck/sip"
)

func main() {
	udp := pcksip.NewUDPTransport()
	udp.Listen(5060, func(msg []byte, addr string) {
		sipmsg, err := pcksip.ParseMessage(msg)

		if err != nil {
			fmt.Println("Error parsing SIP message:", err)
		} else {
			go pcksip.HandlerRequest(addr, sipmsg, udp)
		}
	})
	select {} // 保持程序运行
}
