package main

import (
	"fmt"
	"os"
)

func main() {
	white_list := [4]string{
		"Иван",
		"Пётр",
		"Николай",
		"Сергей",
	}

	var name string
	fmt.Print("Введите имя: ")
	fmt.Scanf("%s", &name)
	for _, white_list_name := range white_list {
		if name == white_list_name {
			fmt.Printf("Hello %s!\n", name)
			os.Exit(0)
		}
	}
	fmt.Printf("Name \"%s\" not allowed. Please, change you name.\n", name)
}
