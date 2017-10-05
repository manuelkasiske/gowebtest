package main

import (
	"github.com/gotschmarcel/goserv"
	"net/http"
		"log"
)

func main() {
	server := goserv.NewServer()
	server.Get("/", func (w http.ResponseWriter, r *http.Request) {
		goserv.WriteString(w, "Welcome Home")
	})
	log.Fatalln(server.Listen(":12345"))
}
