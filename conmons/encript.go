package conmons

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"

	"io"
)

func Encrypting(key []byte, text string) string {

	plainText := []byte(text)

	c, err := aes.NewCipher(key)
	errorHandle(err)

	gcm, err := cipher.NewGCM(c)
	errorHandle(err)

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}

	encryption := gcm.Seal(nonce, nonce, plainText, nil)
	fmt.Println(string(encryption))
	return string(encryption)
}

func Decrypting(key []byte, message string) string {

	ciphertext := []byte(message)

	c, err := aes.NewCipher(key)
	errorHandle(err)

	gcm, err := cipher.NewGCM(c)
	errorHandle(err)

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		fmt.Println(err)
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	errorHandle(err)

	return string(plaintext)
}

// handle this error
func errorHandle(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
