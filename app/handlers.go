package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	mux "github.com/gorilla/mux"
)

// Customer DTO
type Customer struct {
	Name    string `json:"full_name" xml:"name" bson:"full_name"`
	City    string `json:"city" xml:"city" bson:"city"`
	ZipCode string `json:"zip_code" xml:"zip_code" bson:"zip_code"`
}

func greet_handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello World")
}

func getAllCustomers(writer http.ResponseWriter, request *http.Request) {
	customers := []Customer{
		{Name: "John Doe", City: "New York", ZipCode: "10001"},
		{Name: "Jane Doe", City: "New York", ZipCode: "10001"},
		{Name: "John Smith", City: "New York", ZipCode: "10001"},
	}

	if request.Header.Get("Content-Type") == "application/xml" {
		writer.Header().Set("Content-Type", "application/xml")
		xml.NewEncoder(writer).Encode(customers)
		return
	} else if request.Header.Get("Content-Type") == "application/json" {
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(customers)
		return
	}

}

func getCustomer(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	customer_id := vars["customer_id"]
	writer.Write([]byte("Customer ID: " + customer_id))
}

func createCustomer(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Post request received"))
}
