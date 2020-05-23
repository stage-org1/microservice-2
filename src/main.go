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
	fmt.Println("hello world")
	fmt.Fprintf(w, "hello world")
}
