package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"log"
	"os"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

var key = "aljgla.mgh98570fdg;ghjksirl76jnf"

func AesEncrypt(content string) string {
	// 创建加密算法aes
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatalln("Error: NewCipher(%d bytes) = %s", len(key), err)
		os.Exit(-1)
	}

	//加密字符串
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	ciphertext := make([]byte, len(content))
	cfb.XORKeyStream(ciphertext, []byte(content))
	fmt.Printf("%s=>%x\n", content, ciphertext)
	ciphertext = Base64Encode(ciphertext)
	return string(ciphertext)
}

func AesDecrypt(content string) string {
	text := make([]byte, len(content))
	var err error
	text, err = Base64Decode([]byte(content))
	if err != nil {
		log.Fatalln(err)
	}
	// 创建加密算法aes
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatalln("Error: NewCipher(%d bytes) = %s", len(key), err)
		os.Exit(-1)
	}
	// 解密字符串
	decryptText := make([]byte, len(text))
	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	cfbdec.XORKeyStream(decryptText, []byte(text))
	fmt.Printf("%x=>%s\n", text, decryptText)
	return string(decryptText)
}

func Base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

func Base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}
