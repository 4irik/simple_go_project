package main

import "fmt"

func main() {
	var name string
	fmt.Print("Введите имя: ")
	fmt.Scanf("%s", &name)
	fmt.Printf("Hello %s!\n", name)
}
