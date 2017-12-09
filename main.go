package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/LordRahl90/addressapi/handlers"
)

func main() {
	fmt.Println("Starting the addressAPI Server")
	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/create", handlers.Create)

	err := http.ListenAndServe(":9500", nil)

	if err != nil {
		log.Fatal("ListentAndServe", err)
	}
}
