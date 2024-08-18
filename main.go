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
		switch input {
		case "exit":
			os.Exit(0)
		case "print":
			fmt.Println("Список разрешённых имён:")
			for k, name := range whiteList.Names() {
				fmt.Printf("%d. %s\n", k+1, name)
			}
		}
	}
}
