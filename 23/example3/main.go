package main

import "fmt"

func generate(limit int, ch chan<- int) {
	for i := 2; i < limit; i++ {
		fmt.Print(" g", i, "g ")
		ch <- i
	}
	close(ch)
}

func filter(src <-chan int, dst chan<- int, prime int) {
	for i := range src {
		if i%prime != 0 {
			fmt.Print(" *", i, "* ")
			dst <- i
		}
	}
	close(dst)
}

func sieve(limit int) {
	ch := make(chan int)
	go generate(limit, ch)
	for {
		prime, ok := <-ch
		if !ok {
			break
		}
		ch1 := make(chan int)
		fmt.Print(" f", ch, ch1, prime, "f ")
		go filter(ch, ch1, prime)
		ch = ch1
		
		fmt.Print(prime, " ") //glavni ispis
	}
}

func main() {
	sieve(10) // 2 3 5 7 11 13 17 19 23 29 31 37 41 43 ...
}
