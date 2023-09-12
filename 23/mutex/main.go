/*** Example with race condition ***/

package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string) {
	defer wg.Done()
	msg = s
}

func main() {
	msg = "hello world"

	wg.Add(2)
	go updateMessage("Hello universe!")
	go updateMessage("Hello cosmos!")
	wg.Wait()

	fmt.Println(msg)
}

/*** Example with race condition solved with  mutex***/
/*func updateMessage(s string, m *sync.Mutex) {
	defer wg.Done()
	m.Lock()
	msg = s
	m.Unlock()
}

func main() {
	msg = "hello world"

	var mutex sync.Mutex

	wg.Add(2)
	go updateMessage("Hello universe!", &mutex)
	go updateMessage("Hello cosmos!", &mutex)
	wg.Wait()

	fmt.Println(msg)
}*/
