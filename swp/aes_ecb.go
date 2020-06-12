package swp

import (
	"crypto/aes"
	"github.com/andreburgaud/crypt2go/ecb"
	"github.com/andreburgaud/crypt2go/padding"
)

func encryptAESECB(plainText []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	mode := ecb.NewECBEncrypter(block)
	padder := padding.NewPkcs7Padding(mode.BlockSize())
	plainText, err = padder.Pad(plainText)
	if err != nil {
		return nil, err
	}
	ciferText := make([]byte, len(plainText))
	mode.CryptBlocks(ciferText, plainText)
	return ciferText, nil
}

func decryptAESECB(ciferText []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	mode := ecb.NewECBDecrypter(block)
	plainText := make([]byte, len(ciferText))
	mode.CryptBlocks(plainText, ciferText)
	padder := padding.NewPkcs7Padding(mode.BlockSize())
	plainText, err = padder.Unpad(plainText)
	if err != nil {
		return nil, err
	}
	return plainText, nil
}
