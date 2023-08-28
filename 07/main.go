package main

import (
	"bufio"
	"fmt"
	"strings"

	//"io/ioutil"

	//"io"
	"os"
)

func main() {
	for _, fname := range os.Args[1:] {
		file, err := os.Open(fname)
		var lc, wc, cc int // line count, world count, character count
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		/**** first fun io.Copy - kopira sadrzaj fajla na standardni izlaz  ***/

		/*if _, err := io.Copy(os.Stdout, file); err != nil {
			fmt.Fprint(os.Stderr, err)
			continue
		}*/

		/* second fun ioutil.ReadAll - cita ceo sadrzaj fajla u promenljivu data koja je []byte
		/*data, err := ioutil.ReadAll(file)

		if err != nil {
			fmt.Fprint(os.Stderr, err)
			continue
		}

		fmt.Println("The file has", len(data), "bytes")*/

		scan := bufio.NewScanner(file)

		for scan.Scan() {
			s := scan.Text()

			wc += len(strings.Fields(s)) //strings.Fields uzima string i od njega pravi slice stringova, tako da racuna duzinu slice-a
			cc += len(s)
			lc++

		}

		fmt.Printf(" %7d %7d %7d %s\n", lc, wc, cc, fname)

		file.Close()

	}
}
