package main

import (
	"fmt"
	"time"
)

func listenToChan(ch chan int) {
	for {
		i := <-ch
		fmt.Println("Got ", i, " From channel")

		time.Sleep(1 * time.Second)
	}
}

func main() {

	ch := make(chan int, 10) //limit launch goroutine

	defer close(ch)

	go listenToChan(ch)

	for i := 0; i <= 100; i++ {
		fmt.Println("Sending ", i, " to channel")
		ch <- i
		fmt.Println("sent ", i, " to channel")
	}

	fmt.Println("Done!")

}
