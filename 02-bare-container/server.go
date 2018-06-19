package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if res, err := http.Get("http://httpbin.org/get"); err == nil {
			io.Copy(w, res.Body)
			res.Body.Close()
		} else {
			fmt.Fprintf(w, "hello, %s", "world")
		}
	})
	http.ListenAndServe(":8080", nil)
}
