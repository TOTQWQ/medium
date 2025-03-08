package pcksip

import (
	"bufio"
	"fmt"
	"net"
)

type Sip struct {
	conn net.Conn
}

const (
	sipServer = "127.0.0.1:8888"
)

func (s Sip) Client() (*Sip, error) {
	conn, err := net.Dial("udp", sipServer)

	if err != nil {
		return &s, err
	}

	s.conn = conn

	return &s, nil

	// 方法返回前关闭连接
	//defer conn.Close()

}

func (s Sip) Register() (*Sip, error) {
	if s.conn == nil {
		return &s, fmt.Errorf("SIP connection is not established")
	}

	// 生成SIP REGISTER请求
	registerRequest := generateRegisterRequest()

	_, err := s.conn.Write(registerRequest)

	if err != nil {
		fmt.Println("Failed to send SIP REGISTER request:", err)
		return nil, err
	}

	return &s, nil
}

func (s Sip) Read() (string, error) {
	// 读取SIP服务器的响应
	response := make([]byte, 1024)
	_, err := bufio.NewReader(s.conn).Read(response)

	if err != nil {
		fmt.Println("Failed to read SIP response:", err)
		return "", err
	}

	fmt.Println("SIP response received:")
	fmt.Println(string(response))

	return string(response), nil
}

func generateRegisterRequest() []byte {
	// 生成一个简单的SIP REGISTER请求
	// 这里只是一个示例，实际应用中需要根据SIP协议规范生成完整的请求
	request := fmt.Sprintf(`REGISTER sip:%s SIP/2.0
Via: SIP/2.0/UDP 192.168.1.100:5060;branch=z9hG4bK776asdhds
Max-Forwards: 70
From: <sip:alice@example.com>;tag=12345
To: <sip:alice@example.com>
Call-ID: 123456789@192.168.1.100
CSeq: 1 REGISTER
Contact: <sip:alice@192.168.1.100:5060>
Expires: 3600
Content-Length: 0

`, sipServer)

	return []byte(request)
}
