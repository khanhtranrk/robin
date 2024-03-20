package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL)
		http.ServeFile(w, r, "static/index.html")
	})

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("request: favicon icon")
		http.ServeFile(w, r, "assets/favicon.ico")
	})

	http.HandleFunc("/assets/gobin-icon", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("request: gobin icon")
		http.ServeFile(w, r, "assets/gobin-icon.png")
	})

	fmt.Println("Start the server on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
