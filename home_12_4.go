package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	
	"time"
)

func main() {

	var b bytes.Buffer
	var userString string
	for rn:= 1; true; rn++ {
		fmt.Scan(&userString)
		if userString == "выход" {
			break
		}
		t := time.Now().Format("2006-01-02 15:04:05")
		b.Write([]byte(fmt.Sprintf("%d. %s %s \n", rn,t, userString)))

	}
	fileName := "log2.txt"
	if err := ioutil.WriteFile(fileName, b.Bytes(), 0666); err != nil {
		panic(err)
	}
	
	resultBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	
	fmt.Println("Сохраненный лог:")
	fmt.Println(string(resultBytes))



}