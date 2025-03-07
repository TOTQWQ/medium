package pck

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// SIP 方法类型
type Method string

const (
	MethodRegister Method = "REGISTER" // 注册
	MethodInvite   Method = "INVITE"   // 呼叫
	MethodAck      Method = "ACK"      // 应答
	MethodBye      Method = "BYE"      // 挂断
	MethodMessage  Method = "MESSAGE"  // 消息
)

// SIP 消息结构
type Message struct {
	Method    Method
	URI       string
	Version   string
	Headers   map[string]string
	Body      string
	IsRequest bool
}

// 解析 SIP 消息
func ParseMessage(data []byte) (*Message, error) {
	reader := bufio.NewReader(bytes.NewReader(data))
	msg := &Message{Headers: make(map[string]string)}

	// 解析起始行
	startLine, _ := reader.ReadString('\n')
	startLine = strings.TrimSpace(startLine)
	// fmt.Println("起始行：" + startLine)
	parts := strings.Split(startLine, " ")
	if len(parts) < 3 {
		return nil, fmt.Errorf("invalid SIP message")
	}

	if parts[0] == "SIP/2.0" {
		// 响应消息
		msg.IsRequest = false
		msg.URI = parts[1]
		msg.Version = parts[0]
		if len(parts) > 3 {
			msg.Method = Method(parts[2] + " " + parts[3])
		} else {
			msg.Method = Method(parts[2])
		}
	} else {
		// 请求消息
		msg.IsRequest = true
		msg.Method = Method(parts[0])
		msg.URI = parts[1]
		msg.Version = parts[2]
	}

	// 解析头部
	for {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			break
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			msg.Headers[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		}
	}

	// 解析 Body
	body, _ := reader.ReadBytes(0)
	msg.Body = string(body)

	return msg, nil
}

// 生成 SIP 消息
func (m *Message) String() string {
	var buf bytes.Buffer

	if m.IsRequest {
		buf.WriteString(fmt.Sprintf("%s %s %s\r\n", m.Method, m.URI, m.Version))
	} else {
		buf.WriteString(fmt.Sprintf("%s %s %s\r\n", m.Version, m.URI, m.Method))
	}

	for k, v := range m.Headers {
		buf.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
	}
	buf.WriteString("\r\n")
	buf.WriteString(m.Body)

	return buf.String()
}

// 生成 SIP 注册成功响应消息
func NewRegisterSuccessResponse(msg *Message) *Message {
	expires := 3600 // 默认注册有效期为 3600 秒
	if expStr, ok := msg.Headers["Expires"]; ok {
		if exp, err := strconv.Atoi(expStr); err == nil {
			expires = exp
		}
	}

	response := &Message{
		Version: "SIP/2.0",
		URI:     "200",
		Method:  "OK",
		Headers: map[string]string{
			"Via":            msg.Headers["Via"],
			"From":           msg.Headers["From"],
			"To":             msg.Headers["To"],
			"Call-ID":        msg.Headers["Call-ID"],
			"CSeq":           msg.Headers["CSeq"],
			"Contact":        "<sip:34020000002000000001@192.168.2.158:5060>", // 服务器的联系地址
			"Expires":        strconv.Itoa(expires),
			"Content-Length": "0",
		},
	}
	return response
}

// 生成 SIP 注册失败响应消息
func NewRegisterFailResponse(msg *Message) *Message {
	nonce := "1234567890"      // 生成一个随机的 nonce 值
	realm := "sip.example.com" // 认证域

	response := &Message{
		Version: "SIP/2.0",
		URI:     "401 Unauthorized",
		Headers: map[string]string{
			"Via":              msg.Headers["Via"],
			"From":             msg.Headers["From"],
			"To":               msg.Headers["To"] + ";tag=987654321", // 添加 tag 参数
			"Call-ID":          msg.Headers["Call-ID"],
			"CSeq":             msg.Headers["CSeq"],
			"WWW-Authenticate": fmt.Sprintf(`Digest realm="%s", nonce="%s"`, realm, nonce),
			"Content-Length":   "0",
		},
	}
	return response
}
