package main

import (
	"golangweb/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	//closure adalah sebuah function yang disimpan didalam variable
	// aboutHandler := func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("About Page"))
	// }

	//routing
	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/hello", handler.HelloHandler)
	mux.HandleFunc("/deryan", handler.DeryanHandler)
	mux.HandleFunc("/product", handler.ProductHandler)
	mux.HandleFunc("/post-get", handler.PostGet)
	mux.HandleFunc("/form", handler.Form)
	mux.HandleFunc("/process", handler.Process)
	//mux.HandleFunc("/about", aboutHandler)

	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	//route with anomymous func
	// mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Profile Page"))
	// })

	log.Println("Starting web on port 8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
