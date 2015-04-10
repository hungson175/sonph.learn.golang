package main

import (
	"fmt"
	"log"
	"net/http"
)

type String string
type Tokenizer struct {
	Greeting string
	Punct    string
	Who      string
}

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

func (t *Tokenizer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s I am %s", t.Greeting, t.Punct, t.Who)
}
func runWebserver(port int, handler http.Handler) {
	err := http.ListenAndServe(fmt.Sprintf("localhost:%v", port), handler)
	if err != nil {
		log.Fatal(err)
	}

}
func main() {
	stringHandler := String("Hello I am the future Nuts-ProX")
	structHandler := Tokenizer{"Hello", ":", "Son"}
	go runWebserver(3000, stringHandler)
	runWebserver(4000, &structHandler)
}

// package individualfiles

// import (
// 	"fmt"
// 	"net/http"
// )

// //Hello struct
// type Hello struct {
// }

// func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "<html><head><title>Go is so cool </title></head>It is something !<html>")
// }

// func main() {
// 	var h Hello
// 	err := http.ListenAndServe("localhost:3000", h)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
