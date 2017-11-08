package main

import (
	"context"
	"fmt"
	"time"
)

func test() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	out := make(chan int)
	go func() {
		i := 0
		for i < 3 {
			select {
			case out <- i:
				fmt.Println("Send ", i)
				i = i + 1
			case <-ctx.Done():
				fmt.Println("Done")
				fmt.Println(ctx.Err())
				return
			}
		}
		close(out)
	}()

	fmt.Println(<-out)
	fmt.Println("cancal being call")
}

func main() {
	test()
	time.Sleep(5 * time.Second)
}
