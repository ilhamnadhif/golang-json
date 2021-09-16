package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJSON(t *testing.T) {
	jsonString := `{"FirstName":"Muhammad","MiddleName":"Ilham","LastName":"Nadhif","Age":17,"IsMarried":false}`
	jsonBytes := []byte(jsonString)
	//var customer *Customer = &Customer{}
	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)
	fmt.Println(customer.IsMarried)
}
