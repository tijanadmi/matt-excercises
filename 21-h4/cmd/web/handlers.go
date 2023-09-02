package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func (db database) list(w http.ResponseWriter, req *http.Request) {
	

	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) add(w http.ResponseWriter, req *http.Request) {
	

	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	if _, ok := db[item]; ok {
		/*w.WriteHeader(http.StatusBadRequest) // 404

		fmt.Fprintf(w, "duplicate item: %q\n", item)*/
		msg := fmt.Sprintf("duplicate item %q", item)
		http.Error(w,msg,http.StatusBadRequest) //400
		return
	}

	if f64, err := strconv.ParseFloat(price, 32); err != nil {
		/*w.WriteHeader(http.StatusBadRequest) // 400

		fmt.Fprintf(w, "invalid price: %q\n", price)*/
		msg := fmt.Sprintf("invalid price %q", price)
		http.Error(w,msg,http.StatusBadRequest) //400
	} else {
		db[item] = dollars(f64)

		fmt.Fprintf(w, "added %s with price %s\n", item, dollars(f64))
	}
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	

	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	if _, ok := db[item]; !ok {
		/*w.WriteHeader(http.StatusNotFound) // 404

		fmt.Fprintf(w, "no such item: %q\n", item)*/
		msg := fmt.Sprintf("no such item %q", item)
		http.Error(w,msg,http.StatusNotFound) //404
		return
	}

	if f64, err := strconv.ParseFloat(price, 32); err != nil {
		/*w.WriteHeader(http.StatusBadRequest) // 400

		fmt.Fprintf(w, "invalid price: %q\n", price)*/
		msg := fmt.Sprintf("invalid price %q", price)
		http.Error(w,msg,http.StatusBadRequest) //400
	} else {
		db[item] = dollars(f64)

		fmt.Fprintf(w, "new price %s for %s\n", dollars(f64), item)
	}
}

func (db database) fetch(w http.ResponseWriter, req *http.Request) {
	

	item := req.URL.Query().Get("item")

	if _, ok := db[item]; !ok {
		/*w.WriteHeader(http.StatusNotFound) // 404

		fmt.Fprintf(w, "no such item: %q\n", item)*/
		msg := fmt.Sprintf("no such item %q", item)
		http.Error(w,msg,http.StatusNotFound) //404
		return
	}

	fmt.Fprintf(w, "item %s has price %s\n", item, db[item])
}

func (db database) drop(w http.ResponseWriter, req *http.Request) {
	

	item := req.URL.Query().Get("item")

	if _, ok := db[item]; !ok {
		/*w.WriteHeader(http.StatusNotFound) // 404

		fmt.Fprintf(w, "no such item: %q\n", item)*/
		msg := fmt.Sprintf("no such item %q", item)
		http.Error(w,msg,http.StatusNotFound) //404
		return
	}

	delete(db, item)

	fmt.Fprintf(w, "dropped %s\n", item)
}
