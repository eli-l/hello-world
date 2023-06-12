package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	fmt.Println("starting server...")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			panic(err)
		}
	}()
	fmt.Println("server started")
	<-c
	fmt.Println("server stopped")
}
