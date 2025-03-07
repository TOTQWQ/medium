package main

import (
	"fmt"
	"net"
	"strconv"
	"time"

	pcksip "github.com/totqwq/medium/pck/sip"
)

func sendInvite(cameraIP, streamURL string) {
	fmt.Println("----------开始推流--------------")
	conn, err := net.Dial("udp", cameraIP)

	if err != nil {
		fmt.Println("Error connecting to camera:", err)
		return
	}
	defer conn.Close()

	body := fmt.Sprintf(`v=0
o=37070000081118000001 0 0 IN IP4 192.168.2.158
s=Play
c=IN IP4 192.168.2.158
t=0 0
m=video 9000 RTP/AVP 96
a=rtpmap:96 H264/90000`)

	inviteMessage := pcksip.Message{
		Method:    pcksip.MethodInvite,
		URI:       "sip:34020000002000000001@" + cameraIP,
		Version:   "SIP/2.0",
		IsRequest: true,
		Headers: map[string]string{
			"Via":            "SIP/2.0/UDP " + cameraIP + ";branch=z9hG4bK123456",
			"From":           "<sip:37070000081118000001@3402000000>;tag=987654",
			"To":             "<sip:34020000002000000001@192.168.2.188:5060>",
			"Call-ID":        "123456789@192.168.2.158",
			"CSeq":           "1 INVITE",
			"Contact":        "<sip:37070000081118000001@192.168.2.158:5060>",
			"Max-Forwards":   "70",
			"Content-Type":   "application/sdp",
			"Content-Length": strconv.Itoa(len(body)),
		},
		Body: body,
	}

	fmt.Println(inviteMessage.String())

	_, err = conn.Write([]byte(inviteMessage.String()))
	if err != nil {
		fmt.Println("Error sending INVITE:", err)
		return
	}

	fmt.Println("INVITE message sent successfully")
}

func main() {
	cameraIP := "192.168.2.188:5060"                // 摄像头IP
	streamURL := "rtsp://192.168.2.158:9000/stream" // 推流地址

	sendInvite(cameraIP, streamURL)

	// 等待摄像头响应
	time.Sleep(5 * time.Second)
}
