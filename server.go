package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not fount", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not found", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello")
}

func haiHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hai" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not found", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hai")

}

func formData(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform error %v", err)
		return
	}
	fmt.Fprintf(w, "post successful")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "hello %v", name)
	fmt.Fprintf(w, "your address  is %v", address)
}

func main() {

	http.HandleFunc("/form", formData)

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/hai", haiHandler)

	fmt.Println("server starting at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
