package main

import (
	"fmt"
	"log"
	"net/http"
)

type String string

func (h String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, h)
}

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (h Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, h.Greeting + " " + h.Punct + " " + h.Who)
}

func main() {
	http.Handle("/string", String("I'm a frayed knot."));
	http.Handle("/struct", &Struct{Greeting:"Greeting", Punct:"Punct", Who:"Who"});

	// your http.Handle calls here
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
