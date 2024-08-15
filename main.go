package main

import (
	"fmt"
)

func main() {
	var name string
	var whiteList = NewSliceWhiteList()
	whiteList.Add("Иван")
	whiteList.Add("Пётр")
	whiteList.Add("Николай")

loop:
	for {
		fmt.Print("Введите имя: ")
		fmt.Scanf("%s", &name)
		switch true {
		case name == "exit":
			break loop
		case whiteList.IsExist(name):
			fmt.Printf("Hello %s!\n", name)
		default:
			fmt.Printf("Name \"%s\" not allowed. Please, change you name.\n", name)
			break loop
		}
	}
}
