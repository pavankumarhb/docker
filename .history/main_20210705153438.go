package main

import (
	"fmt"
	"net/http"
	"io/ioutil"

)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(w, "welcome to cila family", d)
	})
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request){
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(w, "welcome", d)
	})
	log
}