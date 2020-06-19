package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/about/", about)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatalln(err)
	}

}

func index(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "index.gohtml", nil)
	io.WriteString(res, "This is the index page.")
	HandleError(res, err)
}

func about(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "about.gohtml", nil)
	io.WriteString(res, "this is the about page.")
	HandleError(res, err)
}

func HandleError(res http.ResponseWriter, err error) {
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
