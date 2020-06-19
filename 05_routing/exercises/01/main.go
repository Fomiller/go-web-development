package main

import (
	"io"
	"net/http"
)

type Peach string
type Apple string

func (p Peach) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "This is page is all about Peaches. PEACHES PEACHES PEACHES")
}

func (a Apple) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "This page is all about Apples. APPLES APPLES APPLES")
}

func main() {
	var peach Peach
	var apple Apple

	http.Handle("/peach/", peach)
	http.Handle("/apple/", apple)

	http.ListenAndServe(":3000", nil)
}

// you can also create your own mux but it is not needed by default.
// func main() {
// 	var peach Peach
// 	var apple Apple

// 	mux := http.NewServeMux()
// 	mux.Handle("/peach/", peach)
// 	mux.Handle("/apple/", apple)

// 	// The second value can be named anything mux is just a pointer
// 	http.ListenAndServe(":3000", mux)
// }
