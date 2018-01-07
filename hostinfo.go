package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name, _ := os.Hostname()
		fmt.Fprintf(w, "We are reaching container/host: %s", name)
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}
