package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name" bson:"full_name"`
	City    string `json:"city" xml:"city" bson:"city"`
	ZipCode string `json:"zip_code" xml:"zip_code" bson:"zip_code"`
}

func main() {

	// define routes
	http.HandleFunc("/greet", greet_handler)
	http.HandleFunc("/customers", getAllCustomers)

	// start server
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
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
