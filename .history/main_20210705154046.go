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
		fmt.(w, "welcome to cila family, %s", d)
	})
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request){
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(w, "welcome %s", d)
	})
	http.ListenAndServe(":1234", nil)
}