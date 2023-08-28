package main

import "fmt"

func fib() func() int {
	a, b := 0, 1

	return func() int {
		a, b = b, a+b
		return b
	}
}

func do(d func()){
	d()
}

func main() {
	f, g := fib(), fib()

	fmt.Println(f(), f(), f(), f(), f(), f(), f())
	fmt.Println(g(), g(), g(), g(), g(), g(), g())


	for i := 0; i < 4; i++ {
		v := func() {
			fmt.Printf("%d %p\n", i, &i)
		}
		do(v)
	}
	// bad code
	s := make([]func(), 4)

	for i := 0; i < 4; i++ {
		s[i] = func() {
			// they all point to the same "i"
			fmt.Printf("%d %p\n", i, &i)
		}
	}

	for i := 0; i < 4; i++ {
		s[i]()
	}

	// good code
	s1 := make([]func(), 4)

	for i := 0; i < 4; i++ {
		i2 := i //clouser capture
		s1[i] = func() {
			fmt.Printf("%d %p\n", i2, &i2)
		}
	}

	for j := 0; j < 4; j++ {
		s1[j]()
	}
}