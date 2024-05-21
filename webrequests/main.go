package main

import (
	"fmt"
	"io"
	// "io/ioutil"
	"net/http"
)

const url="http://amazon.com"

func main() {
	fmt.Println("Lco web request")

	res,err:=http.Get(url)
	checkError(err)
	fmt.Printf("Response is of type: %T\n",res)
	defer res.Body.Close()//caller's responsibilty to close the connection
	dataBytes,err:=io.ReadAll(res.Body)
	checkError(err)
	fmt.Println(string(dataBytes))
}

func checkError(err error){
	if err !=nil{
		panic(err)
	}
}