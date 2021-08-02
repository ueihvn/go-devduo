package main

import (
	"log"
	"net/http"

	"github.com/ueihvn/go-devduo/server"
)

func main() {
	s, err := server.NewServer()
	if err != nil {
		log.Fatal(err)
	}
	s.Route()
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", s.Router))
}
