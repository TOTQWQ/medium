package pck

import (
	"fmt"

	"github.com/totqwq/medium/utils"
)

type MessageReceive struct {
	CmdType  string `xml:"CmdType"`
	SN       string `xml:"SN"`
	DeviceID string `xml:"DeviceID"`
	Status   string `xml:"Status"`
	Info     Info   `xml:"Info"`
}

type Info struct {
	// 这里可以添加Info元素下的子元素
}

func HandlerRequest(addr string, msg *Message, udp *UDPTransport) {
	switch msg.Method {
	case MethodInvite: // 呼叫
		fmt.Println("接收到呼叫")
	case MethodAck: // 应答
		fmt.Println("接收到应答")
	case MethodBye: // 挂断
		fmt.Println("接收到挂断请求")
	case MethodRegister: // 注册
		fmt.Println("接受到注册请求")
		// 获取注册成功响应消息
		sendMsg := NewRegisterSuccessResponse(msg)
		// 发送注册成功响应消息
		err := udp.Send(addr, []byte(sendMsg.String()))

		if err != nil {
			fmt.Println("Failed to send response:", err)
		}
	case MethodMessage:
		fmt.Println("接收到消息")
		message := &MessageReceive{}
		if err := utils.XMLDecode([]byte(msg.Body), message); err != nil {
			fmt.Println("Error parsing SIP message:", err)
		} else {
			fmt.Println("body:", message)
		}
	default:
		fmt.Println("Unsupported method:", msg.Method)
	}
}
