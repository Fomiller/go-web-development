package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method        string
		URL           *url.URL
		Submissions   map[string][]string
		Header        http.Header
		Host          string
		ContentLength int64
	}{
		req.Method,
		req.URL,
		req.Form,
		req.Header,
		req.Host,
		req.ContentLength,
	}
	fmt.Printf("***************************\n")
	fmt.Printf("METHOD: %v\n", data.Method)
	fmt.Printf("URL: %v\n", data.URL)
	fmt.Printf("SUBMISSIONS: %v\n", data.Submissions)
	fmt.Printf("HEADER: %v\n", data.Header)
	fmt.Printf("HOST: %v\n", data.Host)
	fmt.Printf("CONTENT LENGTH: %v\n", data.ContentLength)

	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
