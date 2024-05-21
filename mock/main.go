package main 

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main(){
	if len(os.Args)<3{
		log.Fatal("Usage: go run main.go <input_file> <output_file")
	}
	sourceFile:=os.Args[1]
	destintionFile:=os.Args[2]

	err:=copyFile(sourceFile,destintionFile)

	if err != nil{
		log.Fatal(err)
	} 
	fmt.Println("File copied successfully")
}

func copyFile(sourceFile string,destinationFile string)error{
	srcFile,err:=os.Open(sourceFile)
	if err!= nil{
		log.Fatal(err)
	}
	defer srcFile.Close()
	dstFile,err:=os.Create(destinationFile)
	_,err=io.Copy(dstFile,srcFile)
	if err != nil{
		log.Fatal(err)
	}
	return nil
}