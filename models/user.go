package models

import (
	"fmt"
	"github.com/spf13/viper"
	"net/http"

	crunchy "github.com/crunchy-labs/crunchyroll-go/v3"
)

type User struct {
	Email    string
	Password string
	Settings Settings
}

func ExtractUserList() []User {
	var userList []User
	err := viper.UnmarshalKey("Users", &userList)

	if err != nil {
		panic(err)
	}

	fmt.Println(userList)
	return userList
}

func InjectUserList(userList []User) {
	viper.Set("Users", userList)
}

func AddUser(userList []User, email, password string) []User {
	_, err := crunchy.LoginWithCredentials(email, password, crunchy.US, &http.Client{})

	if err != nil {
		panic("Could not sign in. Check if username and password are correct.")
	}

	for idx, _ := range userList {
		if userList[idx].Email == email {
			fmt.Println("User with given email already exists.")
			return userList
		}
	}

	user := User{Email: email, Password: password, Settings: Settings{}}
	fmt.Println("User added.")
	return append(userList, user)
}

func RemoveUser(userList []User, email string) []User {
	for idx, _ := range userList {
		if userList[idx].Email == email {
			fmt.Println("User removed.")
			return removeElement(userList, idx)
		}
	}

	fmt.Println("User with given email does not exist.")
	return userList
}

func removeElement[T any](slice []T, idx int) []T {
	return append(slice[:idx], slice[idx+1:]...)
}
