package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("----------开始推流--------------")
	conn, err := net.Dial("udp", "192.168.2.188:5060")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()

	// 	body := `v=0
	// o=37070000081118000001 0 0 IN IP4 192.168.2.158
	// s=Play
	// c=IN IP4 192.168.2.158
	// t=0 0
	// m=video 43001 TCP/RTP/AVP 96 98 97
	// a=recvonly
	// a=setup:passive
	// a=connection:new
	// a=rtpmap:96 PS/90000
	// a=rtpmap:98 H264/90000
	// a=rtpmap:97 MPEG4/90000
	// a=stream:main
	// a=streamnumber:0
	// y=0000059682`

	// 	to_sip := "sip:37070000081118000001@192.168.2.188:5060"
	// 	contact_sip := "<sip:34020000002000000001@192.168.2.158:5060>"
	// 	inviteMessage := &pcksip.Message{
	// 		Method:    pcksip.MethodInvite,
	// 		URI:       to_sip,
	// 		Version:   "SIP/2.0",
	// 		IsRequest: true,
	// 		Headers: map[string]string{
	// 			"Via":            "SIP/2.0/UDP 192.168.2.158:5060;branch=z9hG4bK0x7f1bcc056800753269171;rport",
	// 			"From":           fmt.Sprintf("%s;tag=957224978", contact_sip),
	// 			"To":             fmt.Sprintf("<%s>", to_sip),
	// 			"Call-ID":        "cfc3d61c-20b5-4bdb-8e33-ff9f8cde911b0",
	// 			"CSeq":           "46311 INVITE",
	// 			"Max-Forwards":   "70",
	// 			"Contact":        contact_sip,
	// 			"User-Agent":     "zdww sip server",
	// 			"Subject":        "37070000081118000001:59682,34020000002000000001:0",
	// 			"Content-Type":   "Application/SDP",
	// 			"Content-Length": strconv.Itoa(len(body)),
	// 		},
	// 		Body: string(body),
	// 	}

	// 	fmt.Println(inviteMessage.String())

	inviteMessage := `INVITE sip:37070000081118000001@192.168.2.188:5060 SIP/2.0
To: <sip:37070000081118000001@192.168.2.188:5060>
From: <sip:34020000002000000001@192.168.2.158:5060>;tag=957224978
Call-ID: cfc3d61c-20b5-4bdb-8e33-ff9f8cde911b0
CSeq: 199946311 INVITE
Max-Forwards: 70
Via: SIP/2.0/UDP 192.168.2.158:5060;branch=z9hG4bK0x7f1bcc056800753269171;rport
Contact: <sip:34020000002000000001@192.168.2.158:5060>
User-Agent: zdww sip server
Subject: 37070000081118000001:59682,34020000002000000001:0
Content-Type: Application/SDP
Content-Length: 294

v=0
o=37070000081118000001 0 0 IN IP4 192.168.2.158
s=Play
c=IN IP4 192.168.2.158
t=0 0
m=video 43001 TCP/RTP/AVP 96 98 97
a=recvonly
a=setup:passive
a=connection:new
a=rtpmap:96 PS/90000
a=rtpmap:98 H264/90000
a=rtpmap:97 MPEG4/90000
a=stream:main
a=streamnumber:0
y=0000059682`

	_, err = conn.Write([]byte(inviteMessage))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("success")
}
