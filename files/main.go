package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "This needs to go in a file - LearnCodeOnline.in"

	file,err:=os.Create("./mylcogofile.txt")
	checkNilErr(err)
	length,err:=io.WriteString(file,content)
	checkNilErr(err)
	fmt.Println("Length is:",length)
	defer file.Close()

	readFile("./mylcogofile.txt")
}

//separte method for reading a file 

func readFile(filename string){
	dataByte,err:=ioutil.ReadFile(filename)
   checkNilErr(err)
	fmt.Println("Text data inside the file is \n ",string(dataByte))
}

func checkNilErr(err error){
	if err!= nil{
		panic(err)
	}
}