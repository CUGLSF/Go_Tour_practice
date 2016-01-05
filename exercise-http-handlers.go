// exercise-http-handlers.go
package main

import (
	"fmt"
	"log"
	"net/http"
)

type String string

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (c String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, string(c))
}
func (p Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s%s%s", string(p.Greeting), string(p.Punct), string(p.Who))
}
func main() {
	// your http.Handle calls here
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
