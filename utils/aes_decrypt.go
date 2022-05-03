package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func Decrypt(text, key, iv string) (string, error) {
	decode_data, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return "", nil
	}
	//生成密码数据块cipher.Block
	block, _ := aes.NewCipher([]byte(key))
	//解密模式
	blockMode := cipher.NewCBCDecrypter(block, []byte(iv))
	//输出到[]byte数组
	origin_data := make([]byte, len(decode_data))
	blockMode.CryptBlocks(origin_data, decode_data)
	//去除填充,并返回
	return string(unpad(origin_data)), nil
}

func unpad(ciphertext []byte) []byte {
	length := len(ciphertext)
	//去掉最后一次的padding
	unpadding := int(ciphertext[length-1])
	return ciphertext[:(length - unpadding)]
}
