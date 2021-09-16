package golang_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamDecoder(t *testing.T) {
	reader, _ := os.Open("sample.json")
	decoder := json.NewDecoder(reader)

	var customer Customer
	decoder.Decode(&customer)
	fmt.Println(customer)
}
