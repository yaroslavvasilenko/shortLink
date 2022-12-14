package main

import (
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	ts, err := template.ParseFiles("./ui/html/home.page.tmpl")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	port := ":4000"
	log.Println("Start server on http://127.0.0.1" + port)

	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}
