package utils

import (
	"fmt"
	"strings"
)

// SDP 结构体表示一个 SDP 会话描述
type SDP struct {
	Version           string
	Origin            Origin
	SessionName       string
	SessionID         string
	SessionVersion    string
	NetworkType       string
	AddressType       string
	Address           string
	Connection        Connection
	Timing            Timing
	MediaDescriptions []MediaDescription
	SSRC              string
}

// Origin 结构体表示 SDP 的原点字段
type Origin struct {
	Username       string
	SessionID      string
	SessionVersion string
	NetworkType    string
	AddressType    string
	Address        string
}

// Connection 结构体表示 SDP 的连接信息字段
type Connection struct {
	NetworkType string
	AddressType string
	Address     string
}

// Timing 结构体表示 SDP 的时间描述字段
type Timing struct {
	Start int64
	Stop  int64
}

// MediaDescription 结构体表示 SDP 的媒体描述字段
type MediaDescription struct {
	MediaType  string
	Port       int
	Protocol   string
	Formats    []string
	Attributes []string
}

// String 方法生成 SDP 字符串
func (s *SDP) String() string {
	var sb strings.Builder

	if s.Version != "" {
		sb.WriteString(fmt.Sprintf("v=%s\r\n", s.Version))
	}

	if s.Origin.Username != "" && s.Origin.SessionID != "" && s.Origin.SessionVersion != "" && s.Origin.Address != "" && s.Origin.AddressType != "" && s.Origin.NetworkType != "" {
		sb.WriteString(fmt.Sprintf("o=%s %s %s %s %s %s\r\n", s.Origin.Username, s.Origin.SessionID, s.Origin.SessionVersion, s.Origin.NetworkType, s.Origin.AddressType, s.Origin.Address))
	}

	if s.SessionName != "" {
		sb.WriteString(fmt.Sprintf("s=%s\r\n", s.SessionName))
	}

	if s.SessionID != "" {
		sb.WriteString(fmt.Sprintf("i=%s\r\n", s.SessionID))
	}

	if s.Connection.NetworkType != "" && s.Connection.AddressType != "" && s.Connection.Address != "" {
		sb.WriteString(fmt.Sprintf("c=%s %s %s\r\n", s.Connection.NetworkType, s.Connection.AddressType, s.Connection.Address))
	}

	sb.WriteString(fmt.Sprintf("t=%d %d\r\n", s.Timing.Start, s.Timing.Stop))

	if s.MediaDescriptions != nil {
		for _, md := range s.MediaDescriptions {
			sb.WriteString(fmt.Sprintf("m=%s %d %s %s\r\n", md.MediaType, md.Port, md.Protocol, strings.Join(md.Formats, " ")))
			if md.Attributes != nil {
				for _, value := range md.Attributes {
					sb.WriteString(fmt.Sprintf("a=%s\r\n", value))
				}
			}
		}
	}

	if s.SSRC != "" {
		sb.WriteString(fmt.Sprintf("y=%s\r\n", s.SSRC))
	}

	return sb.String()
}
