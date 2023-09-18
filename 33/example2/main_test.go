package main

import (
	"testing"
)
var unknown = `{
	"id": 1,
	"name": "bob",
	"addr": {
	"street": "Lazy Lane",
	"city": "Exit",
	"zip": "99999"
	},
	"extra": 21.1
	}`

func TestContains(t *testing.T) {
	var known = []string{
		`{"id": 1}`,
		`{"extra": 21.1}`,
		`{"name": "bob"}`,
		`{"addr": {"street": "Lazy Lane", "city": "Exit"}}`,
	}
	for _, k := range known {
		if err := CheckData([]byte(k), []byte(unknown)); err != nil {
			t.Errorf("invalid: %s (%s)\n", k, err)
		}
	}
}

func TestNotContains(t *testing.T) {
	var known = []string{
			`{"id": 2}`,
			`{"pid": 2}`,
			`{"name": "bobby"}`,
			`{"first": "bob"}`,
			`{"addr": {"street": "Lazy Lane", "city": "Alpha"}}`,
	}
	for _, k := range known {
		if err := CheckData([]byte(k), []byte(unknown)); err == nil {
			t.Errorf("false positive: %s\n", k)
		} else {
			t.Log(err)
		}
	}
}