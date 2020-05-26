package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", reply)
	err:= http.ListenAndServe(":80", nil)
	if (err != nil) {
		fmt.Println(err.Error())
	}
}

func reply(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("time", "0")
	fmt.Fprintf(w, "hello world")
}
