package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	IsMarried  bool
	Hobbies    []string
	Addresses  []Address
}

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Muhammad",
		MiddleName: "Ilham",
		LastName:   "Nadhif",
		Age:        17,
		IsMarried:  false,
	}
	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
