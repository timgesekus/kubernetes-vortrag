package main

import (
    "fmt"
    "net/http"
)

func main() {
	fmt.Println("Starting webserver")
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome to my website - 1.0.0!")
    })

    fs := http.FileServer(http.Dir("static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

	if err := http.ListenAndServe(":9000", nil); err != nil {
		panic(err)
	}
}
