package global

import (
	pcksip "github.com/totqwq/medium/pck/sip"
)

var (
	UDPTransport *pcksip.UDPTransport
	Message      *pcksip.Message
	Addr         string
	TCPTransport *pcksip.TCPTransport
	TCPMessage   *pcksip.Message
)
