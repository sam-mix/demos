package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func main() {

	k1 := []byte("0123456789abcdef")
	r1 := []byte("1234567890abcdef")
	data := []byte("0123456789")
	fmt.Printf("original %x %s\n", data, string(data))

	{
		block, _ := aes.NewCipher(k1)
		stream := cipher.NewCFBEncrypter(block, r1)
		stream.XORKeyStream(data, data)
		fmt.Printf("crypted %x\n", data)
	}
	{
		block, _ := aes.NewCipher(k1)
		stream := cipher.NewCFBDecrypter(block, r1)
		stream.XORKeyStream(data, data)
		fmt.Printf("decrypted %x %s\n", data, string(data))
	}

}
