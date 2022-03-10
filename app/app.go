package app

import (
	"fmt"
	"log"
	"net/http"
)

func Start() {
	http.HandleFunc("/customer", getAllCustomer)
	log.Fatal(http.ListenAndServe("localhost:9000", nil))
	fmt.Println("server started")
}
