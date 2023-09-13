package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m sync.Mutex
	done := make(chan bool)
	go func() {
		m.Lock() // not unlocked!
		//defer m.Unlock() - moramo da dodamo i Unlock inace imamo deadlock
	}()
	go func() {
		time.Sleep(1)
		m.Lock()
		defer m.Unlock()
		done <- true
	}()
	<-done
	fmt.Println("DONE")
}
