package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	gorilla := mux.NewRouter()
	gorilla.HandleFunc("/", IndexHandler)
	gorilla.HandleFunc("/about/", AboutHandler)
	gorilla.HandleFunc("/about/{name}", NameHandler)
	gorilla.HandleFunc("/about/{name}/{food}", FoodHandler)

	err := http.ListenAndServe(":3000", gorilla)
	if err != nil {
		log.Fatalln(err)
	}

}

func IndexHandler(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "index.gohtml", nil)
	io.WriteString(res, "This is the index page.")
	HandleError(res, err)
}

func AboutHandler(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "about.gohtml", nil)
	io.WriteString(res, "this is the about page.")
	HandleError(res, err)
}

func NameHandler(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	err := tpl.ExecuteTemplate(res, "name.gohtml", vars["name"])
	fmt.Println(vars)
	fmt.Println(vars["name"])
	HandleError(res, err)
}

func FoodHandler(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	data := struct {
		Name string
		Food string
	}{
		Name: vars["name"],
		Food: vars["food"],
	}
	err := tpl.ExecuteTemplate(res, "food.gohtml", data)
	fmt.Println(vars)
	fmt.Println(vars["name"])
	HandleError(res, err)
}

func HandleError(res http.ResponseWriter, err error) {
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
