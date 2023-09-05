/** Select example: read two channels **/

package main

import (
	"fmt"
	"time"
)

func main() {
	chans := []chan int{make(chan int), make(chan int)}

	for i := range chans {
		go func(i int, ch chan int) {
			for {
				time.Sleep(time.Duration(i) * time.Second)
				ch <- i
			}
		}(i+1, chans[i])
	}
	for i := 0; i < 12; i++ {
		/***** ovaj slucal prilikom svake iteracije ispisuje ili 1 ili 2 , duplo vise 1 jer u toj go rutini se ceka samo 1sec  ****/
		select {
		case m0 := <-chans[0]:
			fmt.Println("received", m0)
		case m1 := <-chans[1]:
			fmt.Println("received", m1)
		}

		/***** ovaj slucaj ispisuje naizmenicno 1 i 2 u 24 reda ***/
		/*m0 := <-chans[0]
		fmt.Println("received", m0)
		m1 := <-chans[1]
		fmt.Println("received", m1)*/
	}
}
