package main

import (
	"fmt"
	wapi "github.com/codehardt/go-win64api"
)

// userseek - пишет пользователей
func userseek(ch chan<- string) {
	users, err := wapi.ListLoggedInUsers()
	if err != nil {
		fmt.Printf("Ошибка получения списка пользователей: %v\n", err)
		return
	}
	//dbg
	fmt.Println("Пользователи в системе:")
	for _, u := range users {
		user := u.FullUser()
		fmt.Printf("\t%s\n", user)
		ch <- user
	}
}
