package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

type Emps []struct {
	ID             string `json:"id"`
	EmployeeName   string `json:"employee_name"`
	EmployeeSalary string `json:"employee_salary"`
	EmployeeAge    string `json:"employee_age"`
	ProfileImage   string `json:"profile_image"`
}

// https://www.golanglearn.com/http-get-rest-example-using-golang/
func main() {
	e := echo.New()
	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.GET("/employees", func(c echo.Context) error {
		// Build the request
		req, err := http.NewRequest("GET", "http://dummy.restapiexample.com/api/v1/employees", nil)
		if err != nil {
			fmt.Println("Error is req: ", err)
		}

		// create a Client
		client := &http.Client{}

		// Do sends an HTTP request and
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("error in send req: ", err)
		}

		// Defer the closing of the body
		defer resp.Body.Close()

		// Fill the data with the data from the JSON
		var data Emps

		// Use json.Decode for reading streams of JSON data
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			fmt.Println(err)
		}

		return c.JSON(http.StatusOK, data)
	})
}
