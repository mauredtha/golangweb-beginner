package handler

import (
	"golangweb/entity"
	"log"
	"net/http"
	"path"
	"strconv"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// w.Write([]byte("Welcome to My Website"))
	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}

	//type map
	// data := map[string]interface{}{
	// 	"title":   "I'm learning Golang Web",
	// 	"content": "Just do your best, keep learning and remember : practice make perfect",
	// }

	//type struct
	// data := entity.Product{ID: 1, Name: "Fortuner", Price: 420000000, Stock: 3}

	//slice
	data := []entity.Product{
		{ID: 1, Name: "Fortuner", Price: 420000000, Stock: 11},
		{ID: 1, Name: "Xpander", Price: 450000000, Stock: 8},
		{ID: 1, Name: "Pajero Sport", Price: 500000000, Stock: 1},
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Der, Semangat Belajarnya!"))
}

func DeryanHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deryan Everestha Maured"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}
	// w.Write([]byte("Product Page"))

	//fmt.Fprintf(w, "Product page : %d", idNumb)
	data := map[string]interface{}{
		"content": idNumb,
	}

	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}
}

func PostGet(w http.ResponseWriter, r *http.Request) {
	method := r.Method //GET atau POST tergantung cara mengirimkan request ke server

	switch method {
	case "GET":
		w.Write([]byte("Ini adalah GET request"))
	case "POST":
		w.Write([]byte("Ini adalah POST request"))
	default: //PUT, DELETE
		http.Error(w, "Error is happening, keep calm", http.StatusBadRequest)
	}
}

func Form(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles(path.Join("views", "form.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
			return
		}

		return

	}

	http.Error(w, "Error is happening, keep calm", http.StatusBadRequest)
}

func Process(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
			return
		}

		name := r.Form.Get("name")
		message := r.Form.Get("message")

		data := map[string]interface{}{
			"name":    name,
			"message": message,
		}

		tmpl, err := template.ParseFiles(path.Join("views", "result.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
			return
		}
		return
	}

	http.Error(w, "Error is happening, keep calm", http.StatusBadRequest)
}
