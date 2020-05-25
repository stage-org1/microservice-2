package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", reply)
	err:= http.ListenAndServe(":80", nil)
	if (err != nil) {
		fmt.Println(err.Error())
	}
}

func reply(w http.ResponseWriter, r *http.Request) {
	if (rand.Intn(100) < 66) {
		w.Header().Set("time", "0")
		fmt.Fprintf(w, "hello world")
	} else {
		time.Sleep(time.Second * 3)
		w.Header().Set("time", "3")
		fmt.Fprintf(w, "hello world")
	}
}
