package main

import (
	"fmt"
	"image/color"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type IntSlice []int

type ByteCounter int // ovim smo pokazali da je ByteCouter Writer
func (b *ByteCounter) Write(p []byte)(int, error){
	l := len(p)
	*b += ByteCounter(l) //moramo l koji je int da kastujemo u tip koji odgovara b
	return len(p), nil
}

func (is IntSlice) String() string {
	var strs []string
	for _,v := range is {
		strs = append(strs, strconv.Itoa(v))
	}
	return "[" + strings.Join(strs, ";") + "]"
}

type Point struct{
	X, Y float64
}

type Line struct {
	Begin, End Point
}

type Path []Point
func (l Line) Distance() float64{
	return math.Hypot(l.End.X- l.Begin.X, l.End.Y-l.Begin.Y)
}


func (p Path) Distance() (sum float64){
	for i:= 1; i< len(p); i++{
		sum += Line{p[i-1], p[i]}.Distance()
	}
	return sum
}

type Distancer interface{
	Distance() float64
}

func PrintDistance(d Distancer){
	fmt.Println(d.Distance())
}

func (l *Line) ScaleBy(f float64){
	l.End.X += (f-1)*(l.End.X-l.Begin.X)
	l.End.Y += (f-1)*(l.End.Y-l.Begin.Y)
}

type ColoredPoint struct{
	Point
	Color color.RGBA
}

func (p Point) Distance(q Point) float64{
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func main(){

	/*** example 1 **/
	var v IntSlice = []int{1,2,3,4} // v is a Stringer
	var s fmt.Stringer = v// s je interface i mozemo joj dodeliti bilo sta
	

	for i,x := range v {
		fmt.Printf("%d:  %d\n", i, x)
	}

	fmt.Printf("%T %[1]v\n", v)
	fmt.Printf("%T %[1]v\n", s)

	/*** example 2  **/
    var c ByteCounter
	f1, err := os.Open("a.txt") // os.File koji implementira Read i Write metodu
	if err != nil{
	fmt.Fprintf(os.Stderr, "bad a.txt file: %s\n", err)
		os.Exit(-1)
	}

	// check srcFile stats
	fileStat, err := os.Stat("a.txt")
	if err != nil {
		fmt.Print("Failed to check stats for ", "a.txt")
		panic(err)
	}

	// print srcFile stats
	perm := fileStat.Mode().Perm()


	fmt.Printf("File permission before copying %v \n", perm)
	
	
	//f2, err := os.Open("out.txt") // os.File koji implementira Read i Write metodu
	f2 := &c
	if err != nil{
		fmt.Fprintf(os.Stderr, "bad out.txt: %s\n", err)
			os.Exit(-1)
		}

	// check srcFile stats
	fileStat2, err := os.Stat("out.txt")
	if err != nil {
		fmt.Print("Failed to check stats for ", "out.txt")
		panic(err)
	}

	// print srcFile stats
	perm2 := fileStat2.Mode().Perm()
	fmt.Printf("File permission before copying %v \n", perm2)

	n, err := io.Copy(f2, f1) // funkcija io.Copy(dst io.Writer, src io.Reader)
    if err != nil{
		fmt.Fprintf(os.Stderr, "bad copy: %s\n", err)
			os.Exit(-1)
		}
	fmt.Println("copied", n, "bytes")
	fmt.Println(c)


	/*** example 3  **/
	side := Line{Point{1,2}, Point{4,6}}
	perimeter := Path{{1,1}, {5,1},{5,4},{1,1}} // ovde ne moramo da navodimo Point{1,2} zato sto je u pitanju slice []Point

	fmt.Println(side.Distance())
	fmt.Println(perimeter.Distance())
	PrintDistance(side)
	PrintDistance(perimeter)
	side.ScaleBy(3)
	fmt.Println(side.Distance())
	//fmt.Println(Line{Point{1,2}, Point{4,6}}.ScaleBy(2).Distance()) -- ne radi zato sto je nemoguce pozvati Line{Point{1,2}, Point{4,6}}.ScaleBy(2)
	
	p, q := Point{1,1}, ColoredPoint{Point{5,4}, color.RGBA{255,0,0,255}}

	l1 := q.Distance(p)
	l2 := p.Distance(q.Point) // OK; but p.Distance(q) is NOT ALLOWED

	fmt.Println(l1,l2)
}