package main

import (
	"fmt"
	"net/http"
	"io/ioutil"

)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		d, err := ioutil.ReadAll(r.Body)
		if
	})
}