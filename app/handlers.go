package app

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type Customer struct {
	Name    string `json:"name"`
	City    string `json:"city"`
	ZipCode string `json:"zipcode"`
}

func greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Banking Web App!")
}

func getAllcustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "John Doe", City: "New York", ZipCode: "10001"},
		{Name: "Jane Smith", City: "Los Angeles", ZipCode: "90001"},
		{Name: "Mike Johnson", City: "Chicago", ZipCode: "60601"},
		{Name: "Emily Davis", City: "Houston", ZipCode: "77001"},
		{Name: "David Wilson", City: "Miami", ZipCode: "33101"},
	}


	jsonData, err := json.Marshal(customers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}


