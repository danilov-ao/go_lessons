package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
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
	fileName := "ыещкфпу"
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