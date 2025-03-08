package pcksip

import (
	"encoding/json"
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
	is_response := false
	switch msg.Method {
	case MethodInvite: // 呼叫
		fmt.Println("接收到呼叫")
	case MethodAck: // 应答
		fmt.Println("接收到应答")
	case MethodBye: // 挂断
		fmt.Println("接收到挂断请求")
	case MethodRegister: // 注册
		fmt.Println("接受到注册请求")
		fmt.Println("注册信息：" + msg.String())
		is_response = true
	case MethodMessage: // 消息
		fmt.Println("接收到消息")
		message := &MessageReceive{}
		if err := utils.XMLDecode([]byte(msg.Body), message); err != nil {
			// fmt.Println("Error parsing SIP message:", err)
		} else {
			switch message.CmdType {
			case "Keepalive":
				is_response = true
				// fmt.Println(msg.String())
			case "Catalog":
				fmt.Println(msg.String())
				is_response = true
			default:
				fmt.Println("Unsupported message type:", message.CmdType)
			}
		}
	case MethodOK:
		fmt.Println("OK: " + msg.String())
	default:
		// fmt.Println("Unsupported method:", msg.Method)
		fmt.Println("Unsupported method: " + msg.String())
	}

	if is_response {
		// 成功响应消息
		sendMsg := NewRegisterSuccessResponse(msg)
		// fmt.Println(sendMsg.String())
		// 发送成功响应消息
		err := udp.Send(addr, []byte(sendMsg.String()))
		if err != nil {
			fmt.Println("Failed to send response:", err)
		}
	}
}

func (m *MessageReceive) ToJsonString() string {
	// 将 MessageReceive 结构体转换为 JSON 字符串
	jsonData, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	jsonString := string(jsonData)

	// 这里可以根据需要返回一个 *Message 对象
	return jsonString
}
