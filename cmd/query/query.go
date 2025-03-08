package main

import (
	"encoding/xml"
	"fmt"
	"net"
	"strconv"

	pcksip "github.com/totqwq/medium/pck/sip"
	"github.com/totqwq/medium/utils"
)

// Query 结构体表示XML内容
type Query struct {
	XMLName  xml.Name `xml:"Query"`
	CmdType  string   `xml:"CmdType"`
	SN       string   `xml:"SN"`
	DeviceID string   `xml:"DeviceID"`
}

func main() {
	fmt.Println("----------开始推流--------------")
	conn, err := net.Dial("udp", "192.168.2.188:5060")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()

	body, _ := utils.XMLEncode(Query{
		CmdType:  "Catalog",
		SN:       "1708",
		DeviceID: "37070000081118000001",
	})

	fmt.Println("body:", len(body))

	sip := "sip:37070000081118000001@192.168.2.188:5060"

	inviteMessage := &pcksip.Message{
		Method:    pcksip.MethodMessage,
		URI:       sip,
		Version:   "SIP/2.0",
		IsRequest: true,
		Headers: map[string]string{
			"Via":            "SIP/2.0/UDP 192.168.2.158:5060;branch=z9hG4bK0x7fab61e50400385148953;rport",
			"From":           "<sip:34020000002000000001@192.168.2.158:5060>;tag=90650094",
			"To":             fmt.Sprintf("<%s>", sip),
			"Call-ID":        "11ec6335-eaec-4c71-b631-f08d114abae40",
			"CSeq":           "1 MESSAGE",
			"Max-Forwards":   "70",
			"Content-Type":   "Application/MANSCDP+xml",
			"Content-Length": strconv.Itoa(len(body)),
		},
		Body: string(body),
	}

	fmt.Println(inviteMessage.String())

	_, err = conn.Write([]byte(inviteMessage.String()))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("success")
}
