package main

import (
	"fmt"
)

func abc() {
	panic("omg")
}
func main() {
	defer func() {
		if p := recover(); p != nil {
			// what else can you do?
			fmt.Println("recover:", p)
		} else {
			fmt.Println("not panic:")
		}
	}()
	abc()
}
