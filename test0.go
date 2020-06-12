package main

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/openpgp"
	"io/ioutil"
	"os"
)

func main() {
	path := "/home/a/.socofile/057aa60e60fe6504c95960c64a9d0a6a/plain/sample2.txt"
	path2 := "/home/a/.socofile/057aa60e60fe6504c95960c64a9d0a6a/plain/sample2.enc"
	password := []byte("1145141919810")
	
	var (
		enc func() = func() {
			content, _ := ioutil.ReadFile(path)
			
			file, _ := os.OpenFile(path2, os.O_RDWR|os.O_CREATE, 0644)
			
			encrypted, _ := Encrypt(content, password)
			file.Write(encrypted)
			file.Close()
		}
		dec func() = func() {
			content, _ := ioutil.ReadFile(path2)
			decrypted, _ := Decrypt(content, password)
			fmt.Println(decrypted)
			fmt.Println(string(decrypted))
		}
	)
	_, _ = enc, dec
	
	//enc()
	
	dec()
}

func Encrypt(file []byte, passphrase []byte) (encrypted []byte, err error) {
	buffer := bytes.NewBuffer(nil)
	
	plaintext, err := openpgp.SymmetricallyEncrypt(buffer, passphrase, nil, nil)
	if err != nil {
		return nil, err
	}
	_, err = plaintext.Write(file)
	if err != nil {
		return nil, err
	}
	
	plaintext.Close()
	return buffer.Bytes(), nil
}

func Decrypt(file []byte, passphrase []byte) (decrypted []byte, err error) {
	buffer := bytes.NewBuffer(file)
	
	md, err := openpgp.ReadMessage(
		buffer,
		nil,
		func(keys []openpgp.Key, symmetric bool) ([]byte, error) {
			return passphrase, nil
		},
		nil,
	)
	if err != nil {
		return nil, err
	}
	
	decrypted, err = ioutil.ReadAll(md.UnverifiedBody)
	if err != nil {
		return nil, err
	}
	return decrypted, nil
}
