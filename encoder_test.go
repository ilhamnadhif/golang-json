package golang_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("customer_out.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName: "Muhammad",
		MiddleName: "Ilham",
		LastName: "Nadhif",
	}

	encoder.Encode(customer)

	fmt.Println(customer)
}
