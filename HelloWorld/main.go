package main

import (
	"fmt"
	"net/http"
)

func main() {
	//fmt.Println("Hello World")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Easy With Go!")
	})

	http.HandleFunc("/gome", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./home.html")
	})

	http.ListenAndServe(":3000", nil)
}
