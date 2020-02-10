package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	buf := bytes.NewBufferString("one\ntwo\nthree\nfour\n")

	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {

				fmt.Print(line)
				break
			}
			fmt.Println(err)
			break
		}
		fmt.Print(line)
	}
}
