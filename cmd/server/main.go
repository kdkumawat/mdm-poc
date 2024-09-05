package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kdkumawat/mdm-poc/internal/server"
)

func main() {
	port := "3040"

	http.HandleFunc("/policies", server.PoliciesHandler)

	log.Printf("Starting MDM Server on :%v", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
