package main

import (
	"fmt"
	"math"
)

type errFoo struct {
	err error
	path string
}
func (e errFoo) Error() string {
	return fmt.Sprintf("%s: %s", e.path, e.err)
}
func XYZ(a int) *errFoo { return nil }

func TTT(a int) error { return nil }

// Example for method values
type Point struct {
	x,y float64

}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

func (p *Point) DistanceWithPointerReceiver(q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

func main() {

	//err :=  XYZ(1) --- da smo uradili prvo ovo err ne bi bila error interface vec pointer na errFoo a tako ne radimo error/e
	var err error = XYZ(1) // BAD: interface gets a nil concrete ptr
	if err != nil { 
		fmt.Println("oops") 
	}

	// ispisace se "oops" jer error kao interface vise nije nil vec pokazuje na nil

	// u ovom primeru ce se ispisati OK jer f/ja TTT vraca nil pointer na error koji je interface
	var a error =TTT(1)
	if a !=nil{
		fmt.Println("oops")
	}else{
		fmt.Println("OK")
	}

	x:= 2
	g := &x
	fmt.Println(x,g,&x,*g)

	// Example for method values
	
		p := Point{1, 2}
		q := Point{4, 6}
		fmt.Println(p.Distance(q)) 
		distanceFromP := p.Distance // this is a method value
		distanceFromP1 := p.DistanceWithPointerReceiver// this is a method value
		p = Point{2, 3}
		fmt.Println(distanceFromP(q)) // posle promene promenljive p racuna opet istu vrednost, zato sto je stara vrednost zarobljena u deklaraciji funkcije
		fmt.Println(distanceFromP1(q)) // posle promene promenljive p racuna novu vrednos, jer je u funkciji zarobljena adresa, a vrednost na toj adresi se promenila
		
}