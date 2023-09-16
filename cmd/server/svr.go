package main

import (
	"fmt"
	"os"

	"github.com/MrSuds1989/viapolyglot/pkg/model"
)

func main() {
	var user model.User
	fmt.Printf("%#v\n", user)

	user.Id = "abcd"
	user.Name = "Steve"
	user.Email = "test@test.com"
	user.Password = "secret_word!"

	fmt.Printf("%#v\n", user)

	err := user.HashPassword()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%#v\n", user)
}
