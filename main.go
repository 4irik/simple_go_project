package main

import (
	"fmt"
)

func main() {
	whiteList := []string{
		"Иван",
		"Пётр",
		"Николай",
		"Сергей",
	}
	var name string

	fmt.Print("Введите имя: ")
	fmt.Scanf("%s", &name)
	if InWhiteList(name, whiteList) {
		fmt.Printf("Hello %s!\n", name)
	} else {
		fmt.Printf("Name \"%s\" not allowed. Please, change you name.\n", name)
	}
}

func InWhiteList(name string, whileList []string) bool {
	for _, allowedName := range whileList {
		if name == allowedName {
			return true
		}
	}

	return false
}
