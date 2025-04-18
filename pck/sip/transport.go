package pcksip

import (
	"fmt"
	"net"
)

var (
	bufferSize uint16 = 65535 - 20 - 8 // IPv4 max size - IPv4 Header size - UDP Header size
)

// 传输层接口
type Transport interface {
	Send(addr string, msg []byte) error
	Listen(handler func(msg []byte, addr string)) error
}

// UDP 传输实现
type UDPTransport struct {
	conn *net.UDPConn
}

// TCP 传输实现
type TCPTransport struct {
	conn *net.TCPConn
}

func NewUDPTransport() *UDPTransport {
	return &UDPTransport{}
}

func NewTCPTransport() *TCPTransport {
	return &TCPTransport{}
}

func (t *UDPTransport) Send(addr string, msg []byte) error {
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return err
	}
	_, err = t.conn.WriteToUDP(msg, udpAddr)
	return err
}

func (t *UDPTransport) Listen(port int, handler func(msg []byte, addr string)) error {
	// 解析 UDP 地址，监听本地 5060 端口
	udpAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	// 创建 UDP 连接并绑定到指定地址
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		return err
	}

	// 将连接赋值给 t.conn 成员变量
	t.conn = conn

	// 启动一个新的 goroutine 来处理接收到的数据
	go func() {
		buf := make([]byte, bufferSize) // 创建一个 1024 字节的缓冲区
		for {
			// 从连接中读取数据
			n, addr, err := conn.ReadFromUDP(buf)
			if err != nil {
				// 如果读取失败，继续循环（忽略错误）
				continue
			}
			// 调用传入的 handler 处理接收到的消息和发送方地址
			handler(buf[:n], addr.String())
		}
	}()

	// 返回 nil 表示监听成功
	fmt.Printf("-----------开始监听端口：%d------------\r\n", port)
	return nil
}

func (t *TCPTransport) Listen(port int, handler func(data []byte, addr string)) error {
	tcpAddr, err := net.ResolveTCPAddr("udp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	listener, err := net.ListenTCP("udp", tcpAddr)
	if err != nil {
		return err
	}
	defer listener.Close()

	fmt.Printf("-----------开始监听TCP端口：%d------------\r\n", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go func() {
			defer conn.Close()
			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			if err != nil {
				return
			}
			handler(buf[:n], conn.RemoteAddr().String())
		}()
	}
}
