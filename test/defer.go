package main

import (
	"fmt"
)

func main() {
	/*f, err := os.Open("hello.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close() */

	for i := 0; i < 5; i++ {
		defer fmt.Printf("%v ", i)
	}
	// The rest of the program...
}
