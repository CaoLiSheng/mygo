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
	http.HandleFunc("/mobileconfig", func(w http.ResponseWriter, r *http.Request) {
		data, buf := make([]byte, 0), make([]byte, 1024)

		for {
			n, err := r.Body.Read(buf)
			if err == nil {
				for _, v := range (buf[:n]) {
					data = append(data, v)
				}
			} else {
				break;
			}
		}

//		w.WriteHeader(http.StatusMovedPermanently)
//		w.Header().Set("Location", "http://192.168.2.212:10000/enrollment?a=b")
//		w.Write(nil)

		http.Redirect(w, r, "http://192.168.2.212:10000/enrollment", http.StatusFound)

		fmt.Println(String(data))
	});

	// your http.Handle calls here
	log.Fatal(http.ListenAndServe("0.0.0.0:4000", nil))
}
