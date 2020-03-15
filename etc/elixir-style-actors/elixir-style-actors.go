package main

import (
	"log"
)

// https://preslav.me/2020/03/07/elixir-style-actors-in-golang/
func main() {
	in := make(chan message)
	out := make(chan int)
	go newCalculator(0, in, out)

	in <- message{operation: "get"}
	state := <-out
	log.Printf("Current state: %d", state)

	in <- message{operation: "add", value: 100}
	in <- message{operation: "get"}
	state = <-out
	log.Printf("Current state: %d", state)
}

type message struct {
	operation string
	value     int
}

func newCalculator(initialState int, in chan message, out chan int) {
	state := initialState
	for {
		p := <-in
		switch p.operation {
		case "add":
			log.Printf("Adding %d to the current state", p.value)
			state += p.value

		case "get":
			out <- state
		}
	}
}
