package main

import (
	"fmt"
)

func do(b [3]int) int {
	b[0] = 0
	//fmt.Printf("b@ %p\n", b)
	return b[1]
}

func doSlice(b []int) int {
	b[0] = 0
	fmt.Printf("b@ %p\n", b)
	return b[1]
}
func doMap(m1 map[int]int) {
	m1[3] = 0
	fmt.Println("m1", m1)
	fmt.Printf("m1@ %p\n", m1)
	m1 = make(map[int]int)
	m1[4] = 4
	fmt.Println("m1", m1)
	fmt.Printf("m1@ %p\n", m1)
}
func doMapPointer(m1 *map[int]int) {
	(*m1)[3] = 0
	fmt.Println("m1", m1)
	fmt.Println("Pointer of m1 ", &m1)
	*m1 = make(map[int]int)
	(*m1)[4] = 4
	fmt.Println("m1", m1)
	fmt.Println("Pointer of m1 ", &m1)
}
func main() {
	a := [3]int{1, 2, 3}
	c := []int{1, 2, 3}
	m := map[int]int{4: 1, 7: 2, 8: 3}

	fmt.Printf("c@ %p\n", c)

	v := do(a)

	fmt.Println(a, v)

	v1 := doSlice(c)

	fmt.Println(c, v1)

	fmt.Println("Before call fun m", m)
	fmt.Println("Before call fun m@ ", &m)
	doMap(m)
	fmt.Println("m", m)

	mp := map[int]int{4: 1, 7: 2, 8: 3}
	fmt.Println("Before call fun mp", mp)
	fmt.Printf("Before call fun mp@ %p\n", mp)
	doMapPointer(&mp)
	fmt.Println("mp", mp)
}
