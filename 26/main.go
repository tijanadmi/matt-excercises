package main

import (
	"fmt"
	"time"
)

type T struct {
	i byte
	t bool
}

func send(i int, ch chan<- *T) {
	t := &T{i: byte(i)}
	ch <- t
	t.t = true // UNSAFE AT ANY SPEED
}

func main() {
	vs, ch := make([]T, 5), make(chan *T, 2) // change to [5]
	for i := range vs {
		go send(i, ch)
	}

	time.Sleep(1 * time.Second) // all goroutines started

	for i := range vs { // copy quickly!
		vs[i] = *<-ch
	}

	for _, v := range vs { // print later
		fmt.Println(v)
	}
}
