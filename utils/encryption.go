package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"os"
)

func Encrypt(data string) string {
	secret := []byte(os.Getenv("APP_SECRET"))
	iv := []byte(os.Getenv("APP_IV"))
	plainText := []byte(data)

	block, err := aes.NewCipher(secret)
	if err != nil {
		panic(err)
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))

	encryptStream := cipher.NewCTR(block, iv)
	encryptStream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	ivHex := hex.EncodeToString(iv)
	encryptedDataHex := hex.EncodeToString(cipherText)

	return encryptedDataHex[len(ivHex):]
}

func Decrypt(data string) string {
	iv := []byte(os.Getenv("APP_IV"))

	block, err := aes.NewCipher([]byte(os.Getenv("APP_SECRET")))
	if err != nil {
		return ""
	}

	cipherText, err := hex.DecodeString(data)
	if err != nil {
		return ""
	}

	if block.BlockSize() != len(iv) {
		return ""
	}

	ctr := cipher.NewCTR(block, iv)
	plainText := make([]byte, len(cipherText))
	ctr.XORKeyStream(plainText, cipherText)

	return string(plainText)
}
