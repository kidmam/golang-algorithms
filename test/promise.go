package main

import (
	"fmt"
	"time"
)

func myAsyncFunction(s int32) <-chan int32 {
	r := make(chan int32)
	go func() {
		defer close(r)
		// func() core (meaning, the operation to be completed)
		time.Sleep(time.Second * 2)
		r <- s
	}()
	return r
}

func main() {
	firstChannel, secondChannel := myAsyncFunction(2), myAsyncFunction(3)
	first, second := <-firstChannel, <-secondChannel
	// outputs `2, 3` after three seconds
	fmt.Println(first, second)
}
