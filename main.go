package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./Statics"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandle)
	http.HandleFunc("/hello", HelloHandle)
	fmt.Printf("Starting server at port 8000 \n")
	if err := http.ListenAndServe(":8001", nil); err != nil {
		log.Fatal(err)
	}
}
func HelloHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "/GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello!")
}
func formHandle(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err %v :", err)
		return
	}
	fmt.Println(w, "POST request successful")
	name := r.FormValue("Name")
	address := r.FormValue("Address")
	fmt.Fprintf(w, "Name %s \n", name)
	fmt.Fprintf(w, "Address %s \n", address)
}
