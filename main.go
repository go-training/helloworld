package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "<div style: font-size 20px;font-weight:bold;>Hello <p/>Version 3</div>")
    })

    http.ListenAndServe(":8080", nil)
}