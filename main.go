package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ueihvn/go-devduo/server"
)

func main() {
	s, err := server.NewServer()
	if err != nil {
		log.Fatalf("err new server: %v\n", err)
	}

	s.Route()

	fmt.Println("Server run at localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", s.Router))
}
