package main

import (
	"fmt"
	"io"
	
	"os"
)

func main() {
	fmt.Println("welcome to files in golang")
	content := "THis needs to go file"
	file, err := os.Create("./my.text")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println(length)
	defer file.Close()

	readingFile("my.text")
}

func readingFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	dataByte, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dataByte))
}
