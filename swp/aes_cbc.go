package swp

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

func encryptAESCBC(plainText string, key string) (string, error) {
	keyHex, err := hex.DecodeString(key)
	if err != nil {
		return "", err
	}
	plainBytes := []byte(plainText)
	if len(plainText)%aes.BlockSize != 0 {
		plainBytes = append(plainBytes, make([]byte, len(plainText)%aes.BlockSize)...)
	}
	block, err := aes.NewCipher(keyHex)
	if err != nil {
		return "", err
	}
	cipherHex := make([]byte, aes.BlockSize+len(plainBytes))
	iv := cipherHex[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherHex[aes.BlockSize:], plainBytes)
	return fmt.Sprintf("%x", cipherHex), nil
}

func decryptAESCBC(cipherText string, key string) (string, error) {
	keyHex, err := hex.DecodeString(key)
	if err != nil {
		return "", err
	}
	cipherHex, err := hex.DecodeString(cipherText)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(keyHex)
	if err != nil {
		return "", err
	}
	if len(cipherHex) < aes.BlockSize {
		return "", errors.New("Ciphertext too short.")
	}
	iv := cipherHex[:aes.BlockSize]
	cipherHex = cipherHex[aes.BlockSize:]
	if len(cipherHex)%aes.BlockSize != 0 {
		return "", errors.New("Ciphertext is not a multiple of the blocksize")
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherHex, cipherHex)
	return fmt.Sprintf("%s", cipherHex), nil
}
