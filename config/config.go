package config

import (
	"errors"
	"fmt"
	"github.com/leszkolukasz/crunchyroll-notifier/utility"
	"os"
)

func Import(tmpOut **os.File) {
	secret := []byte(utility.GetSecret())
	ensureFileExists("./config/config.json")
	encrypted, err := os.ReadFile("./config/config.json")
	if err != nil {
		panic(err)
	}

	decrypted := utility.Decrypt(encrypted, secret)
	tmp, err := os.CreateTemp("", "tmp*.json")
	if err != nil {
		panic(err)
	}

	_, err = tmp.WriteString(decrypted)
	if err != nil {
		panic(err)
	}

	*tmpOut = tmp
}

func Export(tmp *os.File) {
	secret := []byte(utility.GetSecret())

	stats, err := tmp.Stat()
	if err != nil {
		panic(err)
	}

	decrypted := make([]byte, stats.Size())
	_, err = tmp.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	_, err = tmp.Read(decrypted)
	if err != nil {
		panic(err)
	}

	encrypted := utility.Encrypt(string(decrypted), secret)
	err = os.WriteFile("./config/config.json", encrypted, 666)
	if err != nil {
		panic(err)
	}
}

func GenerateConfiguration() {
	secret := []byte(utility.GetSecret())
	encrypted := utility.Encrypt("{\n\"users\": \"10\"\n}", secret)

	err := os.WriteFile("./config/config.json", encrypted, 666)
	if err != nil {
		panic(err)
	}
}

func ensureFileExists(path string) {
	if _, err := os.Stat(path); !errors.Is(err, os.ErrNotExist) {
		return
	}

	fmt.Println("Configuration file does not exist. Creating...")
	GenerateConfiguration()
	fmt.Println("File created")
}
