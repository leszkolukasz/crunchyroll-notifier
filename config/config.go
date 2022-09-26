package config

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

func Import() {

}

func Export() {

}

func Decrypt(encrypted []byte, secret []byte) string {
	//encrypted, err := os.ReadFile("./config.json")
	//if err != nil {
	//	panic(err)
	//}

	c, err := aes.NewCipher(secret)
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		panic(err)
	}

	nonceSize := gcm.NonceSize()
	if len(encrypted) < nonceSize {
		panic(err)
	}

	nonce, ciphertext := encrypted[:nonceSize], encrypted[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err)
	}

	return string(plaintext)
}

func Encrypt(decrypted string, secret []byte) []byte {
	c, err := aes.NewCipher(secret)
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}

	return gcm.Seal(nonce, nonce, []byte(decrypted), nil)
}
