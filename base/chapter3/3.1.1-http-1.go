package main

import (
	"fmt"
	"net/http"
)

func hiWorld(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "hi, haha")
}

func main() {
  http.HandleFunc("/hi", hiWorld)
  if err := http.ListenAndServe(":8080", nil); err != nil {
    fmt.Println(err)
  }
}
