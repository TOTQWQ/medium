package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("----------查询设备支持推流协议-----------")
	conn, err := net.Dial("udp", "192.168.2.188:5060")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()

	inviteMessage := `OPTIONS sip:37070000081118000001@192.168.2.188:5060 SIP/2.0
Via: SIP/2.0/UDP 192.168.2.158:5060;branch=z9hG4bK123456
Max-Forwards: 70
From: <sip:34020000002000000001@192.168.2.158:5060>;tag=12345
To: <sip:37070000081118000001@192.168.2.188:5060>
Call-ID: cfc3d61c-20b5-4bdb-8w33-ff9f8cde911b0
CSeq: 1 OPTIONS
Contact: <sip:34020000002000000001@192.168.2.158:5060>
Content-Length: 0
`

	_, err = conn.Write([]byte(inviteMessage))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("success")
}
