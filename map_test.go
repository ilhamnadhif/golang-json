package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapEncode(t *testing.T) {
	var product map[string] interface{} = map[string]interface{}{
		"id" : "P0001",
		"name" : "Redmi Note 7",
		"price" : 2700000,
	}
	bytes, err := json.Marshal(product)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
func TestMapDecode(t *testing.T) {
	jsonString := `{"id":"P0001", "name":"Redmi Note 7", "price":2700000}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}

	err := json.Unmarshal(jsonBytes, &result)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
}
