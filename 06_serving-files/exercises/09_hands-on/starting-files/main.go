package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))
	http.HandleFunc("/", surf)
	http.ListenAndServe(":8080", nil)
}

func surf(res http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(res, nil)
	if err != nil {
		log.Fatalln("Template did not execute", err)
	}
}
