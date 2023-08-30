package main

import (
	"fmt"
)

func main() {
	items := [][2]byte{{1, 2}, {3, 4}, {5, 6}}
	a := [][]byte{}

	/*for _, item := range items {
		a = append(a, item[:])
		fmt.Println(a)
	}*/

	for i := range items {
		a = append(a, items[i][:])
	}

	b := [][]byte{}

	for _, item := range items {
		i := make([]byte, len(item))

		copy(i, item[:]) //make unique
		b = append(b, i)
	}

	fmt.Println(items)
	fmt.Println(a)
	fmt.Println(b)
}
