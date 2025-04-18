package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/totqwq/medium/cmd/sip"
	"github.com/totqwq/medium/global"
	pcksip "github.com/totqwq/medium/pck/sip"
)

func main() {
	sip.UDPListen()
	// sip.TCPListen()
	http.HandleFunc("/getMsg", func(w http.ResponseWriter, r *http.Request) {
		request := pcksip.NewQueryRequest()
		fmt.Println(global.Message.Method)
		err := global.UDPTransport.Send(global.Addr, []byte(request))
		if err != nil {
			fmt.Println("Error sending SIP message:", err)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
