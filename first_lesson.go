package main

import (
	"fmt"
	"os"
)

func main() {
	text := "Hello smth"
	file, err := os.Create("hello.txt") //если файл с таким именем уже был то он перезапишется
	if err != nil {
		fmt.Println("Не смогли создать файл")
		return
	}
	defer file.Close()
	file.WriteString(text)
	fmt.Println(file.Name())
}
