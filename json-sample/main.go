package main

import (
	"encoding/json"
	"fmt"
)

// Person is よくあるやつ
type Person struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
}

func main() {
	fmt.Println("string to json ==============")
	// raw string  to  byte array
	bytes := []byte(`[
{
     "id" : 1,
     "name" : "yy_yank",
     "birthday" : "1970/01/01"
},
{
     "id" : 2,
     "name" : "cw_change_word",
     "birthday" : "1970/01/01"
}
]`)
	var persons []Person
	if err := json.Unmarshal(bytes, &persons); err != nil {
		fmt.Println(err)
	}
	for _, p := range persons {
		fmt.Printf("%d : %s :%s\n", p.ID, p.Name, p.Birthday)
	}
	fmt.Println("struct to json string==============")
	var jsonBytes []byte
	var err error
	if jsonBytes, err = json.Marshal(&persons); err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonBytes))
}
