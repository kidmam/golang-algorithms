package main

import "fmt"

func main() {
	options := []int{1, 2, 3}
	funcs := make([]func(), len(options))
	for i, v := range options {
		//
		v := v
		//
		funcs[i] = func() {
			fmt.Println(v)
		}
	}
	for _, f := range funcs {
		f()
	}

	/*for _, v := range options {
		v := v // create a new 'v'.
		go func() {
			fmt.Println(v)
			done <- true
		}()
	}*/
}
