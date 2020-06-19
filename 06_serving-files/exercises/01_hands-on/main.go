package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	// this serves the route requested in the img tag as the source.
	http.HandleFunc("/dog.jpg", doggy)
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Foo Ran.")
}

func dog(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("dog.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(res, "dog.gohtml", nil)
}

func doggy(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "dog.jpg")
}
