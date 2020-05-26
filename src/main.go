package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
	"os"
)

func main() {
	if (rand.Intn(100) < 70) {
		http.HandleFunc("/", reply)
		err:= http.ListenAndServe(":80", nil)
		if (err != nil) {
			fmt.Println(err.Error())
		}
	} else {
		time.Sleep(time.Second*50) //sleep 50 seconds and then crash
		os.Exit(1)
	}
	
}

func reply(w http.ResponseWriter, r *http.Request) {
	if (rand.Intn(100) < 66) {
		w.Header().Set("time", "0")
		fmt.Fprintf(w, "hello world")
	} else {
		time.Sleep(time.Second * 8)
		w.Header().Set("time", "8")
		fmt.Fprintf(w, "hello world")
	}
}
