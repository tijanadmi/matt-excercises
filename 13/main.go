package main

import (
	"fmt"
	"regexp"
	"strings"
)

var uuref = `[[:xdigit:]]{8}-[[:xdigit:]]{4}-[1-5][[:xdigit:]]{3}-[89abAB][[:xdigit:]]{3}-[[:xdigit:]]{12}` // ovako definisano ne obuhvata sve slucajeve. Zato je donja definicija bolja jer ogranicava duzinu UID na tacno odredjenu, gleda tacno odredjena poklapanja
var uure = `^[[:xdigit:]]{8}-[[:xdigit:]]{4}-[1-5][[:xdigit:]]{3}-[89abAB][[:xdigit:]]{3}-[[:xdigit:]]{12}$`
var uu = regexp.MustCompile(uure)
var test = []string{
	"072665ee-a034-4cc3-a2e8-9f1822c4ebbb",
	"072665ee-a034-6cc3-a2e8-9f1822c4ebbb", // wrong version
	"072665ee-a034-4cc3-72e8-9f1822c4ebbb", // wrong format
	"072665ee-a034-4cc3-a2e8-9f1822c4ebbbb",
}

var phre = `\(([[:digit:]]{3})\) ([[:digit:]]{3})-([[:digit:]]{4})`
var pfmt = regexp.MustCompile(phre)

var uure1 = `^(http|https)://([a-zA-Z0-9\-\.]+\.[a-zA-Z]{2,4})` +
	`(?::([0-9]+))?\/?([a-zA-Z0-9\-\._\?\,\'\/\\\+&amp;%\$#\=~]*)$`
var ufmt = regexp.MustCompile(uure1)
var vars = regexp.MustCompile(`(\w+)=(\w+)`)
var test1 = []string{
	"http://matt.com/hello",
	"http://matt.com:8080/hello/",
	"http://matt.com:8080/hello?a=1&b=2",
	"http://matt.com:8080/hello?a=1&b=2&c=3",
}

func main() {
	te := "aba abba abbba"
	re := regexp.MustCompile("b+")
	mm := re.FindAllString(te, -1)
	id := re.FindAllStringIndex(te, -1)

	fmt.Println(re)
	fmt.Println(mm) // [b bb bbb]
	fmt.Println(id) // [[1 2] [5 7] [10 13]]
	for _, d := range id {
		fmt.Println(te[d[0]:d[1]]) // b bb bbb
	}

	te1 := "aba abba abbba"
	re1 := regexp.MustCompile("b+")
	up := re1.ReplaceAllStringFunc(te1, strings.ToUpper)
	fmt.Println(up) // aBa aBBa aBBBa

	for _, t := range test {
		if !uu.MatchString(t) {
			fmt.Println(t, "fails")
		}
	}

	/**** example phone number ***/
	orig := "call me at (214) 514-3232 today"
	match := pfmt.FindStringSubmatch(orig)
	fmt.Printf("%q\n", match)
	if len(match) > 3 {
		fmt.Printf("+1 %s-%s-%s\n", match[1], match[2], match[3])
	}

	intl := pfmt.ReplaceAllString(orig, "+1 ${1}-${2}-${3}")
	fmt.Println(intl)

	/*** URL validation with capture ***/
	for i, t := range test1 {
		match1 := ufmt.FindStringSubmatch(t)
		fmt.Printf("%d: %q\n", i, match1)
	}

	fmt.Println("Second part")

	/*** URL validation with capture part 2 ***/
	for i, t := range test1 {
		match := ufmt.FindStringSubmatch(t)
		fmt.Printf("%d: %q\n", i, match)
		if len(match) > 4 && strings.Contains(match[4], "?") {
			fmt.Printf(" %q\n", vars.FindAllStringSubmatch(match[4], -1))
		}
	}

}
