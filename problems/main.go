package main

import (
	"io"
	"log"
	"net"
	"time"
)

// "fmt"
// "time"

// func fun(s string){
// 	for i:=0;i<3;i++{
// 		fmt.Println(s)
// 		time.Sleep(1 *time.Millisecond)
// 	}
// }

func main(){
	//exercise -1 
	//Direct call 
	// fun("direct call")
	// go fun("goroutine-1")

    // go func(){
	// 	fun("goroutine-2")
	// }()


	// fmt.Println("Wait for goroutines...")
	// time.Sleep(100*time.Millisecond)
	// fmt.Println("done...")

    
	//excercise-client-server

	listener,err:=net.Listen("tcp","localhost:8000")
	if err != nil{
		log.Fatal(err)
	}
	for{
	conn,err:=	listener.Accept()
	if err!=nil{
		continue
	}
	handleConn(conn)
	}
	

}

func handleConn(c net.Conn){
	defer c.Close()
	for{
		_,err :=io.WriteString(c,"response from servver\n")
		if err !=nil{
			return 
		}
		time.Sleep(time.Second)
	}
}