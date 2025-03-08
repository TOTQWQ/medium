package sip

import (
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
