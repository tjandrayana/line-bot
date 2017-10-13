package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("Halo")
	http.HandleFunc("/index/", viewHandler)
	http.ListenAndServe(":1000", nil)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {

	a := []byte("test : aaa")

	w.Header().Set("Content-Type", "application/json")
	w.Write(a)

}
