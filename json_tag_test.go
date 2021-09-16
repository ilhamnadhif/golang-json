package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string	`json:"name"`
	ImageURL string `json:"image_url"`
}

func TestJSONTag(t *testing.T) {
	product := Product{
		Id: "P0001",
		Name: "Redmi Note 7",
		ImageURL: "https://google.com/",
	}
	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {
	//jsonString := `{"ID":"P0001","NAME":"Redmi Note 7","IMAGE_URL":"https://google.com/"}`
	jsonString := `{"id":"P0001","name":"Redmi Note 7","image_url":"https://google.com/"}`
	jsonBytes := []byte(jsonString)

	var product Product
	err := json.Unmarshal(jsonBytes, &product)
	if err != nil {
		panic(err)
	}

	fmt.Println(product)
	fmt.Println(product.Name)
}