/*** Problem filozofa koji jedu, ako imamo vise od jednog Mutax-a uvek ih zakljucavamo i otkljucavamo u odredjenom redosledu  ****/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m1, m2 sync.Mutex
	done := make(chan bool)

	fmt.Println("START")
	go func() {
		m1.Lock()
		defer m1.Unlock()
		time.Sleep(1)
		m2.Lock()
		defer m2.Unlock()
		fmt.Println("SIGNAL 1")
		done <- true
	}()
	go func() {
		/**** ovakav redosled Mutax-a ne prolazi izaziva deadlock ***/
		/*m2.Lock()
		defer m2.Unlock()
		time.Sleep(1)
		m1.Lock()
		defer m1.Unlock()*/

		m1.Lock()
		defer m1.Unlock()
		time.Sleep(1)
		m2.Lock()
		defer m2.Unlock()
		fmt.Println("SIGNAL 2")
		done <- true
	}()
	<-done
	fmt.Println("DONE after first done")
	<-done
	fmt.Println("DONE after second done")
}
