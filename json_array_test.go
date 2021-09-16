package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArrayEncode(t *testing.T) {

	customer := Customer{
		FirstName:  "Muhammad",
		MiddleName: "Ilham",
		LastName:   "Nadhif",
		//Age:        17,
		//IsMarried:  false,
		Hobbies: []string{"Game", "Reading", "Coding"},
	}
	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Muhammad","MiddleName":"Ilham","LastName":"Nadhif","Age":0,"IsMarried":false,"Hobbies":["Game","Reading","Coding"]}`
	jsonBytes := []byte(jsonString)

	var customer *Customer = &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
	fmt.Println(customer.Hobbies[1])
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		MiddleName: "Ilham",
		Addresses: []Address{
			{
				Street:     "Jalan 1",
				Country:    "Indonesia",
				PostalCode: "88888",
			},
			{
				Street:     "Jalan 2",
				Country:    "Indonesia",
				PostalCode: "88888",
			},
		},
	}
	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"","MiddleName":"Ilham","LastName":"","Age":0,"IsMarried":false,"Hobbies":null,"Addresses":[{"Street":"Jalan 1","Country":"Indonesia","PostalCode":"88888"},{"Street":"Jalan 2","Country":"Indonesia","PostalCode":"88888"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.Addresses)
	fmt.Println(customer.Addresses[0].Street)
}
func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Jalan 1","Country":"Indonesia","PostalCode":"88888"},{"Street":"Jalan 2","Country":"Indonesia","PostalCode":"88888"}]`
	jsonBytes := []byte(jsonString)

	addresses := []Address{}
	err := json.Unmarshal(jsonBytes, &addresses)
	if err != nil {
		panic(err)
	}
	fmt.Println(addresses)
	fmt.Println(addresses[0].Street)
}

func TestOnlyJSONArrayComplex(t *testing.T) {
	addresses := []Address{
		{
			Street:     "Jalan 1",
			Country:    "Indonesia",
			PostalCode: "88888",
		},
		{
			Street:     "Jalan 2",
			Country:    "Indonesia",
			PostalCode: "88888",
		},
	}
	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))
}
