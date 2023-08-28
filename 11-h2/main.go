package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var raw = `
<!DOCTYPE html>
<html>
  <body>
    <h1>My First Heading</h1>
      <p>My first paragraph.</p>
      <p>HTML <a href="https://www.w3schools.com/html/html_images.asp">images</a> are defined with the img tag:</p>
      <img src="xxx.jpg" width="104" height="142">
  </body>
</html>`

func visit(n *html.Node, pwords, ppics *int){

	if n.Type == html.TextNode{
		*pwords += len(strings.Fields(n.Data))
		
	} else if n.Type == html.ElementNode && n.Data == "img"{
		*ppics++
	}

	for c:= n.FirstChild; c != nil; c= c.NextSibling {
		visit(c, pwords,ppics)
	}
}

func countWordsAndImages(doc *html.Node) (int, int) {
	var words, pics int

	visit(doc, &words, &pics)

	return words, pics
}

func main() {
	doc, err := html.Parse(bytes.NewReader([]byte(raw)))

	if err != nil {
		fmt.Fprintf(os.Stderr, "parse failed: %s\n", err)
		os.Exit(-1)
	}
	

	words, pics := countWordsAndImages(doc)

	fmt.Printf("%d words and %d images\n", words, pics)


	s := `<p>Links:</p>
	      <ul>
		  <li>
		  <a href="foo">Foo</a>
		  <li><a href="/bar/baz">BarBaz</a></ul>`
	doc1, err := html.Parse(strings.NewReader(s))
	if err != nil {
		log.Fatal(err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			
			for _, a := range n.Attr {
				if a.Key == "href" {
					fmt.Println(a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc1)
}