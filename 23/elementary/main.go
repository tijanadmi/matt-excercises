package main

import (
	"fmt"
	"sync"
	"time"
)

type Book struct {
	Title  string
	Author string
}

var cache = make(map[int]Book)

func queryCache(id int) (Book, bool) {
	b, ok := cache[id]
	return b, ok
}

func queryDatabase(id int, b Book) (Book, bool) {
	cache[id] = b
	return b, true
}
func main() {

	/*for i := 0; i < 10; i++ {

		go func(id int) {
			fmt.Println(id)

		}(i)
		go func(id int) {
			fmt.Println(id)
			time.Sleep(3 * time.Microsecond)

		}(i)

		time.Sleep(150 * time.Microsecond)
	}*/
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(2)
		go func(id int, wg *sync.WaitGroup) {
			fmt.Println(id)
			wg.Done()
		}(i, wg)
		go func(id int, wg *sync.WaitGroup) {
			fmt.Println(id)
			time.Sleep(3 * time.Microsecond)
			wg.Done()
		}(i, wg)

		time.Sleep(150 * time.Microsecond)
	}
	wg.Wait()

	/*cache[1] = Book{"Proces", "Franc Kafka"}
	cache[2] = Book{"Pakao", "Dante Aligijeri"}
	cache[3] = Book{"1984", "Dzordz Orvel"}
	cache[4] = Book{"Duga setnja do slobode", "Nelson Mandela"}
	for i := 0; i < 10; i++ {
		go func(id int) {
			if b, ok := queryCache(id); ok {
				fmt.Println(b.Title, b.Author)
			}
		}(i)
		go func(id int) {
			if b, ok := queryCache(id); !ok {
				if b, ok := queryDatabase(id, b); ok {
					fmt.Println(b.Title, b.Author)
				}
			}
		}(i)
		time.Sleep(150 * time.Microsecond)
	}*/

}
