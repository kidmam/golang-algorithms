package main

import (
	"fmt"
	"time"
)

func myAsyncFunction() <-chan int32 {
	r := make(chan int32)
	go func() {
		defer close(r)
		// func() core (meaning, the operation to be completed)
		time.Sleep(time.Second * 2)
		r <- 2
	}()
	return r
}

func main() {
	r := <-myAsyncFunction()
	// outputs `2` after two seconds
	fmt.Println(r)
}
