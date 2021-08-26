package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	fmt.Println("введите имя пользователя")
	var userName string
	var password string
	var age int
	fmt.Scan(&userName)
	fmt.Println("введите пароль")
	fmt.Scan(&password)
	fmt.Println("введите возраст")
	fmt.Scan(&age)

	file, err := os.Create("creadenials.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
// Важный момент - объявляется буффер
// для чего нужен интерфейс (со слов лектора это контракт поведения) интерфейс io
// Дословно если чтото имеет такие же свойства используя интерфейс его можно вызывать у другого метода; в GO интерфейсы не явные
// пока аксиома у Buffera есть метод write.String
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("Ваш логин:%s",userName))
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf("Ваш пароль:%s",password))
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf("Ваш возраст:%d",age))
	b.WriteString("\n")
	_, err = file.Write(b.Bytes()) // что означает нижнее подчеркивание ?
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
}
