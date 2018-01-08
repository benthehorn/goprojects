package main

import (
	"fmt"
	"log"
	"net/http"
)

func heyWorld(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hello World!\n")
}

func main(){
	port := 8080

	http.HandleFunc("/", heyWorld)

	log.Printf("Server starting on port %v\n", port)
	http.ListenAndServe(fmt.Sprintf(":%v",  port), nil)
}