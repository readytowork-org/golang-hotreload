package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Golang Devs"))
	})

	PORT := "5000"

	fmt.Println("Server running on port:", PORT)

	http.ListenAndServe(":"+PORT, nil)

}
