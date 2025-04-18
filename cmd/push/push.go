package main

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"

	pcksip "github.com/totqwq/medium/pck/sip"
	"github.com/totqwq/medium/utils"
)

var (
	passageway = "34020000001318000105"
	device_ip  = "192.168.2.105:5060"
	sip_ip     = "192.168.2.115:5060"
	sip_id     = "34020000002000000001"
	domain     = "3402000000"
	sip_user   = "34020000001188000009"
)

func main() {
	fmt.Println("----------开始推流--------------")
	conn, err := net.Dial("udp", device_ip)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()

	ssrc := "0000059682"

	body := &utils.SDP{
		Version: "0",
		Origin: utils.Origin{
			Username:       sip_user,
			SessionID:      "0",
			SessionVersion: "0",
			NetworkType:    "IN",
			AddressType:    "IP4",
			Address:        sip_ip,
		},
		SessionName: "Play",
		Connection: utils.Connection{
			NetworkType: "IN",
			AddressType: "IP4",
			Address:     sip_ip,
		},
		Timing: utils.Timing{
			Start: 0,
			Stop:  0,
		},
		MediaDescriptions: []utils.MediaDescription{
			{
				MediaType: "video",
				Port:      32014,
				Protocol:  "TCP/RTP/AVP",
				Formats:   []string{"96", "97", "98", "99"},
				Attributes: []string{
					"recvonly", "rtpmap:96 PS/90000", "rtpmap:98 H264/90000",
					"rtpmap:97 MPEG4/90000", "rtpmap:99 H265/90000",
					"setup:passive", "connection:new",
				},
			},
		},
		SSRC: ssrc,
	}

	inviteMessage := &pcksip.Message{
		Method:    pcksip.MethodInvite,
		URI:       fmt.Sprintf("sip:%s@%s", passageway, device_ip),
		Version:   "SIP/2.0",
		IsRequest: true,
		Headers: map[string]string{
			"Via":            fmt.Sprintf("SIP/2.0/UDP %s;rport;branch=z9hG4bK%s", sip_ip, strconv.Itoa(rand.Intn(1000000))),
			"From":           fmt.Sprintf("<sip:%s@%s>;tag=%s", sip_id, domain, strconv.Itoa(rand.Intn(1000000))),
			"To":             fmt.Sprintf("<%s@%s>", passageway, device_ip),
			"Call-ID":        fmt.Sprintf("%s@192.168.2.158", strconv.Itoa(rand.Intn(1000000))),
			"CSeq":           "20 INVITE",
			"Max-Forwards":   "70",
			"Contact":        fmt.Sprintf("<sip:%s@%s>", sip_id, sip_ip),
			"User-Agent":     "zdww sip server",
			"Subject":        fmt.Sprintf("%s:%s,%s:%s", passageway, ssrc, sip_id, "0"),
			"Content-Type":   "Application/SDP",
			"Content-Length": strconv.Itoa(len(body.String())),
		},
		Body: body.String(),
	}

	fmt.Println(inviteMessage.String())

	_, err = conn.Write([]byte(inviteMessage.String()))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("success")
}
