package androidlib

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

var key = "aljgla.mgh98570fdg;ghjksirl76jnf"

func AesDecrypt(content string) string {
	text := make([]byte, len(content))
	var err error
	text, err = Base64Decode([]byte(content))
	if err != nil {
		return content
	}
	// 创建加密算法aes
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return content
	}
	// 解密字符串
	decryptText := make([]byte, len(text))
	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	cfbdec.XORKeyStream(decryptText, []byte(text))
	return string(decryptText)
}

func Base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}
