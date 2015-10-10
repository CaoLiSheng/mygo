package main

import "net/http"

func main() {
	http.ListenAndServe(":10000", http.FileServer(http.Dir("/Users/caols/git/mygo/public/")))
}
