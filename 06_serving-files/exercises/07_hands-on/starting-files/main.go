package main

import (
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("./public/"))))
	http.HandleFunc("/", dog)
	http.ListenAndServe(":8080", nil)
}

func dog(res http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(res, nil)
	if err != nil {
		log.Fatalln("Template did not execute", err)
	}
}
