package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", reply)
	err:= http.ListenAndServe(":80", nil)
	if err != nil {
		os.Exit(1)
	}
}


func reply(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}