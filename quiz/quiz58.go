package main

import (
	"fmt"
)

func main() {
	var a *int
	b := 5
	c := 12
	a = &c
	b = c
	*a = 14
	fmt.Println(b)
}
