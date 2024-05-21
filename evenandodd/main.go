package main

import (
	"fmt"
	"time"
)

func main() {
	evenCh := make(chan int)
	oddCh := make(chan int)

	go generateNumbers(evenCh, oddCh)

	for i := 0; i < 20; i++ {
		select {
		case even := <-evenCh:
			fmt.Printf("%d is even\n", even)
		case odd := <-oddCh:
			fmt.Printf("%d is odd\n", odd)
		}
	}
}

func generateNumbers(evenCh, oddCh chan int) {
	for i := 1; ; i++ {
		time.Sleep(time.Second) // Simulate generation time
		if i%2 == 0 {
			evenCh <- i
		} else {
			oddCh <- i
		}
	}
}
