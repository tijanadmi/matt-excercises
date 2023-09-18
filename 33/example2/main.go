package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func matchNum(key string, exp float64, data map[string]interface{}) bool {
	if v, ok := data[key]; ok {
		if val, ok := v.(float64); ok && val == exp {
			return true
		}
	}
	return false
}

func matchString(key, exp string, data map[string]interface{}) bool {
	// is it in the map?
	if v, ok := data[key]; ok {
		// is it a string, and does it match?
		if val, ok := v.(string); ok && strings.EqualFold(val, exp) {
			return true
		}
	}
	return false
}

func contains(exp, data map[string]interface{}) error {
	for k, v := range exp {
		fmt.Println(k,v)
		
		switch x := v.(type) {
			case float64:
				if !matchNum(k, x, data) {
					return fmt.Errorf("%s unmatched (%d)", k, int(x))
				}
			case string:
				if !matchString(k, x, data) {
					return fmt.Errorf("%s unmatched (%s)", k, x)
				}
		case map[string]interface{}:
				if val, ok := data[k]; !ok {
					return fmt.Errorf("%s missing in data", k)
				} else if unk, ok := val.(map[string]interface{}); ok {
					if err := contains(x, unk); err != nil {
						return fmt.Errorf("%s unmatched (%+v): %s", k, x, err)
					}
				} else {
					return fmt.Errorf("%s wrong in data (%#v)", k, val)
				}
		}
	}
	return nil
}

func CheckData(want, got []byte) error {

	var w, g map[string]interface{}
	if err := json.Unmarshal(want, &w); err != nil {
		return err
	}
	if err := json.Unmarshal(got, &g); err != nil {
		return err
	}
	return contains(w, g)
}



func main(){
	var got = `{
		"id": 1.1,
		"name": "bob",
		"addr": {
		"street": "Lazy Lane",
		"city": "Exit",
		"zip": "99999"
		},
		"extra": 21.1
		}`
		var want =  `{"id": 1.1}`
		var w, g map[string]interface{}
	if err := json.Unmarshal([]byte(want), &w); err != nil {
		fmt.Println("Greska 1 ", err)
	}
	if err := json.Unmarshal([]byte(got), &g); err != nil {
		fmt.Println("Greska 2 ",err)
	}

		if err := contains(w, g); err != nil {
			fmt.Println("ne sadrzi",err)
		}else{
			fmt.Println("sadrzi")
		}
}
				