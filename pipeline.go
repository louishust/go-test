package main

import (
	"fmt"
	"time"
)

func main() {
	out := make(chan int)
	go func() {
		i := 0
		for i < 3 {
			out <- i
			fmt.Println("Send ", i)
			i = i + 1
		}
		close(out)
	}()

	fmt.Println(<-out)
	time.Sleep(1 * time.Minute)
}
