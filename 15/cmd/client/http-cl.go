package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const url = "https://jsonplaceholder.typicode.com"

type todo struct {
	ID        int    `json:"id"` //ime polja mora pocinjati velikim slovom ako zelimo da nam se dekodira u JSON-u
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	//resp, err := http.Get("http://localhost:8080/" + os.Args[1])

	resp, err := http.Get(url + "/todos/1")

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}

		//fmt.Println(string(body))

		var item todo

		err = json.Unmarshal(body, &item)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}
		fmt.Printf("%#v\n", item)
	}

}
