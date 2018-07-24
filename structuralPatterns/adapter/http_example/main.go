package main

import (
	"fmt"
	"log"
	"net/http"
)

type MyServer struct {
	Msg string
}

func (m *MyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func main() {
	server := &MyServer{
		Msg: "Hello, World",
	}

	http.Handle("/", server)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
