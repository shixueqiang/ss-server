package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"log"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

var key = "aljgla.mgh98570fdg;ghjksirl76jnf"

func AesEncrypt(content string) string {
	// 创建加密算法aes
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatalf("Error: NewCipher(%d bytes) = %s\n", len(key), err)
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
		log.Fatalf("Error: NewCipher(%d bytes) = %s\n", len(key), err)
	}
	// 解密字符串
	decryptText := make([]byte, len(text))
	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	cfbdec.XORKeyStream(decryptText, []byte(text))
	fmt.Printf("%x=>%s\n", text, decryptText)
	return string(decryptText)
}

func AesEncrypt1(origData []byte) string {
	var cryptedKey = []byte(key)
	block, err := aes.NewCipher(cryptedKey)
	if err != nil {
		log.Fatalf("Error: NewCipher(%d bytes) = %s\n", len(key), err)
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, cryptedKey[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return string(Base64Encode(crypted))
}

func AesDecrypt1(crypted []byte) string {
	var cryptedKey = []byte(key)
	block, err := aes.NewCipher(cryptedKey)
	if err != nil {
		log.Fatalf("Error: NewCipher(%d bytes) = %s\n", len(key), err)
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, cryptedKey[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return string(origData)
}

func Base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

func Base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
