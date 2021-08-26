package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

func main(){
	var b bytes.Buffer

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(101)
	fmt.Println("Введите число от 1 до 100")
	b.WriteString("Введите число от 1 до 100\n")
	for {
		var answer int
		for {
			_,_=fmt.Scan(&answer)
			b.WriteString(fmt.Sprintf("Введено число %d \n",answer))
			if answer < 1 || answer > 100 {
				fmt.Println("Число должно быть в диапазоне от 1 до 100")
				b.WriteString("Число должно быть в диапазоне от 1 до 100\n")
			} else {
				break
			}
		}
		if answer == n {
			fmt.Println("Число угаданно")
			b.WriteString("Число угаданно \n")
			break
		} else if answer < n  {
			fmt.Println("Загаданное число больше")
			b.WriteString("Загаданное число больше \n")
		} else {
			fmt.Println("Загаданное число меньше")
			b.WriteString("Загаданное число меньше \n")

		}
	}
	fileName := "log.txt"
	if err := ioutil.WriteFile(fileName, b.Bytes(), 0666); err != nil {
		panic(err)
	}
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	resultBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Println("Сохраненный лог:")
	fmt.Println(string(resultBytes))



}
