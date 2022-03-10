package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type Customer struct {
	Name    string `json:"name"`
	City    string `json:"city"`
	ZipCode int    `json:"zip_code"`
}

func getAllCustomer(w http.ResponseWriter, r *http.Request) {
	customer := []Customer{
		{Name: "sheetal", City: "delhi", ZipCode: 322001},
		{Name: "Ashish", City: "udaipur", ZipCode: 321233},
	}
	if r.Header.Get("Content-Type") == "application/json" {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
	} else {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customer)
	}
}
