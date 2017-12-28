package main 

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

func main() {
	

	http.HandleFunc("/", func(w http.ResponseWriter, r*http.Request) {
		fmt.Fprint(w, time.Now().Format(time.RFC850),  html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r*http.Request) {
		fmt.Fprint(w, "Hi")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}