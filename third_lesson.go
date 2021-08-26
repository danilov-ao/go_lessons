package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("log.txt")
	if err != nil {
		fmt.Println("Ошибка открытия файла", err)
		return
	}
	defer f.Close()
	/*buf := make ([]byte,128 )
	if _, err := io.ReadFull(f,buf); err != nil {
		fmt.Println("Не смогли прочитать последовательность бфтов из файла ", err)
		return
	}
	fmt.Printf("%s\n", buf)
*/
	buf := make ([]byte,150 )
	_, err = f.Read(buf)
	if err != nil {
		fmt.Println("Не смогли прочитать последовательность бфтов из файла ", err)
		return
	}
	fmt.Println(string(buf))
}

