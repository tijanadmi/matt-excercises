/***** primer deadlock-a, program ceka na nesto sto se nikada nece desiti *****/

package main

import (
	"fmt"
)

func main() {
	ch := make(chan bool)
	go func(ok bool) {
		fmt.Println("STARTED")
		if ok {
			ch <- ok
		}
	}(false)
	<-ch //program ceka na nesto sto se nikada nece desiti
	fmt.Println("DONE")
}
