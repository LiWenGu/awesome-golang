package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	server := http.Server{
		Addr:      ":8848",
		Handler:   nil,
		TLSConfig: nil,
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("HandlerFunc called!\n")
		writer.Write([]byte("Hello World"))
	})
	err := server.ListenAndServeTLS("./src/单向认证/server.crt", "./src/单向认证/server.key")

	if err != nil {
		log.Fatal(err)
	}
}
