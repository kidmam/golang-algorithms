package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Student struct {
	Name string
	Id   int
}

func main() {
	students := []Student{
		{
			Name: "Sumit",
			Id:   1,
		},
		{
			Name: "Amit",
			Id:   2,
		},
	}

	j, err := json.Marshal(students)
	if err != nil {
		log.Fatalf("Error in Marshalling, %s", err)
	}
	fmt.Printf("%s\n", j) // [{"Name":"Sumit","Id":1},{"Name":"Amit","Id":2}]
}
