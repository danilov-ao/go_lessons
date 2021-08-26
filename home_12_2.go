package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("storage.txt")
	if err != nil {
		fmt.Println("Не смогли открыть файл", err)
		return
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		fmt.Printf("Рамер файла не считан")
	}
	//Создаем некоторую послдовательность byte
	buf := make([]byte, fi.Size())
	if _, err := io.ReadFull(f, buf); err != nil {
		fmt.Println("Не смогли прочитать последовательность файлов из byte", err)
		return
	}
	fmt.Printf("%s \n", buf)
}