package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main(){
	fmt.Println("Hello")
	// PerformGetRequest()
	PerformPostRequest()

	
}


func PerformGetRequest(){
	const myurl="https://jsonplaceholder.typicode.com/posts"
	res,err:=http.Get(myurl)
    if err != nil{
		panic(err)
	}
	defer res.Body.Close()
	fmt.Println("status code: ",res.StatusCode)
	fmt.Println("content length is:", res.ContentLength)
	var responseString strings.Builder
	content,_:=ioutil.ReadAll(res.Body)
	byteCount,_:=responseString.Write( content)
	fmt.Println(byteCount)
	// fmt.Println(responseString.String())
	fmt.Println(responseString)


	
	// fmt.Println(string(content))

}

func PerformPostRequest(){
	const myurl="https://jsonplaceholder.typicode.com/posts" 

	//fake json payload

	requestBody:=strings.NewReader(`
	{
		"title": "foo",
		"body": "bar",
		"userId": "101"
	  }
	

	`)
	res,err:=http.Post(myurl, "application/json",requestBody)
	if err != nil{
		panic(err)
	}

	defer res.Body.Close()
	content,err:=ioutil.ReadAll(res.Body)
	if err != nil{
		panic(err)
	}
	fmt.Println(string(content))
	fmt.Println(res.StatusCode)
	

}