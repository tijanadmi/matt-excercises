package main

import (
	"fmt"
	"log"
	"net/http"
)

//var nextID =0
var nextID = make(chan int)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>You got %v<h1>", <-nextID)
	// unsafe - data race
	//nextID ++
}

func counter() {
	for i := 0; ; i++ {
		nextID <- i
	}
}
func main() {
	go counter()
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
