package controller

import (
	"encoding/json"
	"fmt"
	"github.com/kidmam/golang-algorithms/etc/go-cache/helper"
	"net/http"
	_ "reflect"
)

var SaveCache = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var emp helper.Emp
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&emp)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error saving cache")
		return
	}
	helper.SetCache("emp_data", emp)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "successfully! data has been saved in cache")
})

var GetCache = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	data, found := helper.GetCache("emp_data")
	if found {
		fmt.Println(data)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "successfully! data has been saved in cache")
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error! Not found key into cache")
		return
	}
})
