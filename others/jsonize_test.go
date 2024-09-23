package others

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func jsonize() {
	person := Person{
		Name: "Rama",
		Age:  22,
	}

	personJson, err := json.Marshal(person)
	fmt.Println(personJson)
	fmt.Printf("%v", err)
}

func TestJsonize(t *testing.T) {
	jsonize()
}

func unjsonize() {
	personByte := `{"name":"armadillo","age":19}`
	var person Person
	if err := json.Unmarshal([]byte(personByte), &person); err != nil {
		log.Fatal(err)
	}
	fmt.Println(person)
}

func TestUnjsonize(t *testing.T) {
	unjsonize()
}
