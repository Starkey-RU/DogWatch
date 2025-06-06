package main

import (
	"fmt"
	wapi "github.com/codehardt/go-win64api"
)

// userseek finds all currently logged in users
func userseek(ch chan<- string) {
	users, err := wapi.ListLoggedInUsers()
	if err != nil {
		fmt.Printf("Error capturing logged in users: %v\n", err)
		return
	}
	
	fmt.Println("Users found:")
	for _, u := range users {
		user := u.FullUser()
		fmt.Printf("\t%s\n", user)
		ch <- user
	}
}
