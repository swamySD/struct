package main

import (
	"fmt"
	"net/http"
	"time"
)

func main(){
	fmt.Println("Welcome")
	links:=[]string{
		"http://google.com",
		"http://facebook.com",
		"http://amazon.com",
		"http://udemy.com",
		"http://stackoverflow.com",
		"http://instagram.com",
		
	}
	c:=make(chan string)
	for _,link :=range links{
	 go	checkLink(link, c)
	}
  for l:=range c{
	
	go func(link string){
	time.Sleep(5 *time.Second)
	 checkLink(link,c)
       }(l)
	
	}
}

func checkLink(link string ,c chan string){
	_,err:=http.Get(link)
	checkErr(err,link,c)
	fmt.Println(link,"is up!")
	c <- link
}


func checkErr(err error,link string, c chan string){
	if err!= nil{
fmt.Println(link,"Might be down")
 c <- link
		panic(err)
	}
}



// func readNumber()[]int{
// 	var numbers []int
// 	var num int 
// 	for{
// 		_,err:=fmt.Scan(&num)
// 		if err != nil{
// 			break
// 		}
// 		numbers=append(numbers, num)
// 	}
	
// 	return numbers
	
// }

