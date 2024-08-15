package main

import (
	"fmt"
	"os"
)

func main() {
	var input string
	var whiteList = NewSliceWhiteList()
	whiteList.Add("Иван")
	whiteList.Add("Пётр")
	whiteList.Add("Николай")

	fmt.Print("Введите имя: ")
	fmt.Scanf("%s", &input)
	if !whiteList.IsExist(input) {
		fmt.Printf("Name \"%s\" not allowed. Please, change you name.\n", input)
		os.Exit(0)
	}

	for {
		fmt.Print("Введите комманду: ")
		fmt.Scanf("%s", &input)
		if input == "exit" {
			os.Exit(0)
		}
	}
}
