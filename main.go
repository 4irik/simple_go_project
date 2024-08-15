package main

import (
	"fmt"
)

func main() {
	var name string
	var whiteList = new(SliceWhiteList)
	whiteList.Add("Иван")
	whiteList.Add("Пётр")
	whiteList.Add("Николай")

	fmt.Print("Введите имя: ")
	fmt.Scanf("%s", &name)
	if whiteList.IsExist(name) {
		fmt.Printf("Hello %s!\n", name)
	} else {
		fmt.Printf("Name \"%s\" not allowed. Please, change you name.\n", name)
	}
}
