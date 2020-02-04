package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/sirupsen/logrus"
	"net/http"
)

type New_Emp struct {
	Name   string `json:"name"`
	Salary string `json:"salary"`
	Age    string `json:"age"`
}

// https://www.golanglearn.com/json-post-payload-example-in-golang/
func main() {
	routes := mux.NewRouter()
	routes.HandleFunc("/add_employee", AddEmployee).Methods("POST")
}

func AddEmployee(w http.ResponseWriter, r *http.Request) {
	var emp New_Emp
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&emp)

	jsonValue, _ := json.Marshal(emp)
	fmt.Printf("%+v\n", emp)
	u := bytes.NewReader(jsonValue)

	req, err := http.NewRequest("POST", "http://dummy.restapiexample.com/api/v1/create", u)
	if err != nil {
		fmt.Println("Error is req: ", err)
	}

	req.Header.Set("Content-Type", "application/json")
	// create a Client
	client := &http.Client{}

	// Do sends an HTTP request and
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error in send req: ", err.Error())
		w.WriteHeader(400)
		//w.Write(err)
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var data New_Emp
	res, err := json.Unmarshal(resp.Body, &data)
	w.Write(data)
}
