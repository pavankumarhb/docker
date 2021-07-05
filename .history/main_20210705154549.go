package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, "welcome to cila family, %s", d)
	})
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request){
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, "welcome %s", d)
	})
	http.ListenAndServe(":8000", nil)
}