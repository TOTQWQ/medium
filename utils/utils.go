package utils

import (
	"bytes"
	"encoding/xml"
	"io"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// XMLDecode 解析XML
func XMLDecode(data []byte, v interface{}) error {
	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel
	return decoder.Decode(v)
}

// GBK 转 UTF-8
func GbkToUtf8(data []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(data), simplifiedchinese.GBK.NewDecoder())
	d, e := io.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

// 结构体转XML
func XMLEncode(data any) ([]byte, error) {
	data, err := xml.Marshal(data)
	if err != nil {
		return nil, err
	}

	byteData, ok := data.([]byte)

	if !ok {
		return nil, err
	}

	return byteData, nil
}

// 可封装通用XML生成函数（放置于utils包中）
func XMLEncodeWithHeader(v interface{}, header string) ([]byte, error) {
	data, err := XMLEncode(v)
	if err != nil {
		return nil, err
	}
	return append([]byte(header), data...), nil
}
