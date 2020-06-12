package file

import (
	"bytes"
	"golang.org/x/crypto/openpgp"
	"io/ioutil"
)

// Symmetrically encryption & decryption for bytes data
// using OpenPGP package

// Modified from https://gist.github.com/jyap808/8250124

func SymmEncrypt(file []byte, passphrase []byte) (encrypted []byte, err error) {
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

func SymmDecrypt(file []byte, passphrase []byte) (decrypted []byte, err error) {
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
