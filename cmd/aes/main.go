package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"strconv"
)

// AES加密
// iv为空则采用ECB模式，否则采用CBC模式
func AesEncrypt(value, secretKey, iv string) (string, error) {
	if value == "" {
		return "", nil
	}

	//根据秘钥生成16位的秘钥切片
	keyBytes := make([]byte, aes.BlockSize)
	copy(keyBytes, []byte(secretKey))
	//获取block
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	blocksize := block.BlockSize()
	valueBytes := []byte(value)

	//填充
	fillsize := blocksize - len(valueBytes)%blocksize
	repeat := bytes.Repeat([]byte{byte(fillsize)}, fillsize)
	valueBytes = append(valueBytes, repeat...)

	result := make([]byte, len(valueBytes))

	//加密
	if iv == "" {
		temp := result
		for len(valueBytes) > 0 {
			block.Encrypt(temp, valueBytes[:blocksize])
			valueBytes = valueBytes[blocksize:]
			temp = temp[blocksize:]
		}
	} else {
		//向量切片
		ivBytes := make([]byte, aes.BlockSize)
		copy(ivBytes, []byte(iv))

		encrypter := cipher.NewCBCEncrypter(block, ivBytes)
		encrypter.CryptBlocks(result, valueBytes)
	}

	//以hex格式数值输出
	encryptText := fmt.Sprintf("%x", result)
	return encryptText, nil
}

// AES解密
// iv为空则采用ECB模式，否则采用CBC模式
func AesDecrypt(value, secretKey, iv string) (string, error) {
	if value == "" {
		return "", nil
	}

	//根据秘钥生成8位的秘钥切片
	keyBytes := make([]byte, aes.BlockSize)
	copy(keyBytes, []byte(secretKey))
	//获取block
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	//将hex格式数据转换为byte切片
	valueBytes := []byte(value)
	var encryptedData = make([]byte, len(valueBytes)/2)
	for i := 0; i < len(encryptedData); i++ {
		b, err := strconv.ParseInt(value[i*2:i*2+2], 16, 10)
		if err != nil {
			return "", err
		}
		encryptedData[i] = byte(b)
	}

	result := make([]byte, len(encryptedData))

	if iv == "" {
		blocksize := block.BlockSize()
		temp := result
		for len(encryptedData) > 0 {
			block.Decrypt(temp, encryptedData[:blocksize])
			encryptedData = encryptedData[blocksize:]
			temp = temp[blocksize:]
		}
	} else {
		//向量切片
		ivBytes := make([]byte, aes.BlockSize)
		copy(ivBytes, []byte(iv))

		//解密
		blockMode := cipher.NewCBCDecrypter(block, ivBytes)
		blockMode.CryptBlocks(result, encryptedData)
	}

	//取消填充
	unpadding := int(result[len(result)-1])
	result = result[:(len(result) - unpadding)]
	return string(result), nil
}

func main() {
	text := "上山打老虎"
	key := "123456"
	iv := "abcdefg"

	encryptText1, _ := AesEncrypt(text, key, iv)
	fmt.Printf("【%s】经过【AES-CBC】加密后：%s\n", text, encryptText1)

	decryptText1, _ := AesDecrypt(encryptText1, key, iv)
	fmt.Printf("【%s】经过【AES-CBC】解密后：%s\n", encryptText1, decryptText1)

	encryptText2, _ := AesEncrypt(text, key, "")
	fmt.Printf("【%s】经过【AES-ECB】加密后：%s\n", text, encryptText2)

	decryptText2, _ := AesDecrypt(encryptText2, key, "")
	fmt.Printf("【%s】经过【AES-ECB】解密后：%s\n", encryptText2, decryptText2)
}
