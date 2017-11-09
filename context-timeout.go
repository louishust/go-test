package main

import (
	"context"
	"fmt"
	"time"
)

func test() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
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
	for i := 1; i <= 10; {
		time.Sleep(1 * time.Second)
		fmt.Println("wait for ", i, " sec ")
		i = i + 1
	}
}

func main() {
	test()
}
