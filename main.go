package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var input string
	whiteList := NewSliceWhiteList()

	filePath := flag.String("file", "", "Путь до файла с именами")
	flag.Parse()

	if *filePath == "" {
		fmt.Println("Путь до файла не задан. Загружается стандартный набор имён")
		setDefaultNames(whiteList)
	} else {
		fmt.Println("Чтение имён из файла:", *filePath)
		bs, err := os.ReadFile(*filePath)
		if err != nil {
			panic(err)
		}
		for _, name := range strings.Split(string(bs), "\n") {
			whiteList.Add(name)
		}
	}

	for {
		fmt.Print("Введите имя: ")
		fmt.Scanf("%s", &input)
		if !whiteList.IsExist(input) {
			fmt.Printf("Name \"%s\" not allowed. Please, change you name.\n", input)
			os.Exit(0)
		}

		autorityActionLoop(whiteList)
	}
}

func setDefaultNames(whiteList WhiteList) {
	whiteList.Add("Иван")
}

func autorityActionLoop(whiteList WhiteList) {
	var input string

loop:
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
		case "add":
			fmt.Scanf("%s", &input)
			if !whiteList.IsExist(input) {
				whiteList.Add(input)
				fmt.Printf("OK. Имя \"%s\" добавлено в список\n", input)
			} else {
				fmt.Printf("Warn. Имя \"%s\" уже в списке\n", input)
			}
		case "delete":
			fmt.Scanf("%s", &input)
			if whiteList.IsExist(input) {
				whiteList.Delete(input)
				fmt.Printf("OK. Имя \"%s\" идалено из списка\n", input)
			} else {
				fmt.Printf("Warn. Имя \"%s\" отсутствует в списке\n", input)
			}
		case "logout":
			break loop
		default:
			fmt.Printf("Команда \"%s\" не распознана\n", input)
		}
		input = ""
	}
}
