package sip

import (
	"fmt"

	"github.com/totqwq/medium/global"
	pcksip "github.com/totqwq/medium/pck/sip"
)

func UDPListen() {
	global.UDPTransport = pcksip.NewUDPTransport()
	global.UDPTransport.Listen(5060, func(msg []byte, addr string) {
		sipmsg, err := pcksip.ParseMessage(msg)

		if err != nil {
			// fmt.Println("Error parsing SIP message:", err)
		} else {
			global.Message = sipmsg
			global.Addr = addr
			pcksip.HandlerRequest(addr, sipmsg, global.UDPTransport)
		}
	})
}

func TCPListen() {
	global.TCPTransport = pcksip.NewTCPTransport()
	global.TCPTransport.Listen(9000, func(msg []byte, addr string) {
		fmt.Println("tcp:", string(msg))
	})
}
