package main

import (
	"fmt"
	"path/filepath"
	"sort"
)

type Pair struct{
	Path string
	Hash string
}

func (p Pair) String() string{
	return fmt.Sprintf("Hash of %s is %s", p.Path, p.Hash)
}

type PairWithLength struct{
	Pair
	Length int
}

func (p PairWithLength) String() string{
	return fmt.Sprintf("Hash of %s is %s; length %d", p.Path, p.Hash, p.Length)
}

/*func Filename(p Pair) string {
	return filepath.Base(p.Path)
}*/
func (p Pair) Filename() string {
	return filepath.Base(p.Path)
}

type Filenamer interface{
	Filename() string
}

type Fizgig struct {
	*PairWithLength
	Broken bool
}

/**** iz primera za sortiranje  ***/
type Organ struct {
	Name string
	Weight int
	}
type Organs []Organ
func (s Organs) Len() int { return len(s) }
func (s Organs) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

type ByName struct{ 
	Organs 
}

func (s ByName) Less(i, j int) bool {
	return s.Organs[i].Name < s.Organs[j].Name
}
type ByWeight struct{ 
	Organs 
}
func (s ByWeight) Less(i, j int) bool {
	return s.Organs[i].Weight < s.Organs[j].Weight
}

///example with Stack -- make the nil value useful
type StringStack struct {
	data []string // "zero" value ready-to-use -- data je definisano malim slovom, tako da se ne vidi van paketa (u slucaju da ovaj kod stavimo u paket koji cemo importovati u drugi i ovako definisan slice je nil
} 

func (s *StringStack) Push(x string) {
	s.data = append(s.data, x)
}
func (s *StringStack) Pop() string {
	if l := len(s.data); l > 0 {
		t := s.data[l-1]
		s.data = s.data[:l-1]
		return t
	}
	panic("pop from empty stack")
}

type IntList struct {
	Value int
	Tail *IntList
}

// Sum returns the sum of the list elements.
func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}

func main(){
	p:= Pair{"/usr", "0xfdfe"}
	pl := PairWithLength{Pair{"/usr/lib", "0xdead"}, 133}
	var fn Filenamer = PairWithLength{Pair{"/usr/lib", "0xdead"}, 133}  // ova dodela vazi jer embeddet type implementira interface
	fmt.Println(p)
	//fmt.Println(Filename(p)) // prvi slicaj
	
	fmt.Println(fn)

	fmt.Println(pl)
	// prvi slucaj fmt.Println(pl.Pair.Filename) -- ne mozemo napisati fmt.Println(Filename(pl)) jer f-ja Filename kao argument prima polje tipa Pair
	fmt.Println(pl.Filename()) 

	fg := Fizgig{
		&PairWithLength{Pair{"/usr", "0xfdfe"}, 121},
		false,
	}
	fmt.Println(fg) // implementira metodu String embeddet type-a

	//sorting with StringSlice

	entries := []string{"charlie", "able", "dog", "baker"}
	sort.Sort(sort.StringSlice(entries))
	fmt.Println(entries) 

	// primer sortiranja sa organima
	s := []Organ{
		{"brain", 1340}, {"heart", 290},
		{"liver", 1494}, {"pancreas", 131},
		{"spleen", 162},
		}

	sort.Sort(ByWeight{s}) // pancreas first
	fmt.Println(s)
	sort.Sort(ByName{s}) // brain first
	fmt.Println(s)

	// defined in the sort package
	// type StringSlice []string
	entries1 := []string{"charlie", "able", "dog", "baker"}
	fmt.Println(entries1) 
	sort.Sort(sort.Reverse(sort.StringSlice(entries1))) // type casting
	fmt.Println(entries1) // [dog charlie baker able]
	
}