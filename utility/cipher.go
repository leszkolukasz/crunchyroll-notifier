package utility

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

func Decrypt(encrypted []byte, secret []byte) string {
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

func GetSecret() string {
	secret, ok := os.LookupEnv("CRUNCHY_SECRET")

	if !ok {
		panic(fmt.Errorf("no CRUCNHY_SECRET in environment variables"))
	}

	return secret
}
