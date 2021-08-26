package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	file, err := os.Create("storage.txt")
	if err := os.Chmod("storage.txt", 0444); err != nil {
		fmt.Println(err)
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var userString string
	for rn := 1; true; rn++ {
		fmt.Scan(&userString)
		if userString == "выход" {
			break
		}
		t := time.Now().Format("2006-01-02 15:04:05")
		file.Write([]byte(fmt.Sprintf("%d. %s %s \n", rn, t, userString)))

	}

}
