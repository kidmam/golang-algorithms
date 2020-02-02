package main

import (
	"github.com/kabukky/httpscerts"
	"io"
	"log"
	"net/http"
)

// https://www.golanglearn.com/webserver-over-https-using-golang-and-tls/
func main() {

	http.HandleFunc("/home", ExampleHandler)

	log.Println("** Service Started on Port 8081 **")

	// Use ListenAndServeTLS() instead of ListenAndServe() which accepts two extra parameters.
	// We need to specify both the certificate file and the key file (which we've named
	// https-server.crt and https-server.key).
	// Check if the cert files are available.
	err := httpscerts.Check("cert.pem", "key.pem")
	// If they are not available, generate new ones.
	if err != nil {
		err = httpscerts.Generate("cert.pem", "key.pem", "localhost:8081")
		if err != nil {
			log.Fatal("Error: Couldn't create https certs.")
		}
	}
	err1 := http.ListenAndServeTLS(":8081", "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatal(err1)
	}
}

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, `{"status":"ok"}`)
}
