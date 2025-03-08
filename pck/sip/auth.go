package pcksip

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

// AuthParams 存储认证参数
type AuthParams struct {
	Username   string
	Realm      string
	Nonce      string
	URI        string
	Response   string
	Algorithm  string
	Qop        string
	NonceCount string
	Cnonce     string
}

// CalculateResponse 计算 Digest Authentication 的 Response
func (a *AuthParams) CalculateResponse(method, password string) string {
	ha1 := a.calculateHA1(password)
	ha2 := a.calculateHA2(method)
	response := a.calculateResponse(ha1, ha2)
	return response
}

// calculateHA1 计算 HA1
func (a *AuthParams) calculateHA1(password string) string {
	data := fmt.Sprintf("%s:%s:%s", a.Username, a.Realm, password)
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}

// calculateHA2 计算 HA2
func (a *AuthParams) calculateHA2(method string) string {
	data := fmt.Sprintf("%s:%s", method, a.URI)
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}

// calculateResponse 计算最终的 Response
func (a *AuthParams) calculateResponse(ha1, ha2 string) string {
	data := fmt.Sprintf("%s:%s:%s", ha1, a.Nonce, ha2)
	if a.Qop == "auth" {
		data = fmt.Sprintf("%s:%s:%s:%s:%s:%s", ha1, a.Nonce, a.NonceCount, a.Cnonce, a.Qop, ha2)
	}
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}

// ParseAuthHeader 解析 WWW-Authenticate 或 Authorization 头
func ParseAuthHeader(header string) map[string]string {
	params := make(map[string]string)
	parts := strings.Split(header, ",")
	for _, part := range parts {
		kv := strings.SplitN(strings.TrimSpace(part), "=", 2)
		if len(kv) == 2 {
			params[kv[0]] = strings.Trim(kv[1], `"`)
		}
	}
	return params
}
