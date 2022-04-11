package Serialization

import (
	"encoding/json"
	"fmt"
	"testing"
)

// type Person struct {
// 	Name    string `json:name`
// 	Gender  string `json:gender`
// 	Age     int    `json:age`
// 	address string `json:address`
// }
type Person struct {
	Name    string
	Gender  string
	Age     int
	address string
}

func TestSerialization(t *testing.T) {
	p1 := Person{
		"jack",
		"boy",
		18,
		"New York",
	}

	p, _ := json.Marshal(p1)
	fmt.Println(p)

	var p2 Person
	json.Unmarshal(p, &p2)
	fmt.Println(p2)
}
