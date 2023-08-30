package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type todo struct {
	UserID    int    `json:"userID"`
	ID        int    `json:"id"` //ime polja mora pocinjati velikim slovom ako zelimo da nam se dekodira u JSON-u
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var form = `
<h1>Todo #{{.ID}}</h1>
<div>{{printf "User %d" .UserID}}</div>
<div>{{printf "%s (completed: %t)" .Title  .Completed}}</div>`

const base = "https://jsonplaceholder.typicode.com/"

func handler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "hello, world! from %s\n", r.URL.Path[1:])

	var item todo
	resp, err := http.Get(base + r.URL.Path[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
	defer resp.Body.Close()
	//body, _ := ioutil.ReadAll(resp.Body)
	err = json.NewDecoder(resp.Body).Decode(&item) // umesto ioutil.ReadAll i json.Unmarshal(body, &item)  -- cita Body u vidu JSON-a i smesta u strukturu -- mozemo proslediti razlicite source-ove io.Reader-u i Dedode ce ih decodirati citajuci bajtove direktno sa Reader-a
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
	//_ = json.Unmarshal(body, &item)
	tmpl := template.New("mine")
	tmpl.Parse(form)
	tmpl.Execute(w, item)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
