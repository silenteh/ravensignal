package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// api := NewApi("127.0.0.1", "8080")
	// api.Start()

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8181", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
