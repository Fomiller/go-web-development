package main

import (
	"html/template"
	"log"
	"os"
)

type hotels []hotel

type hotel struct {
	Name, Address, City string
	Region              region
	Zip                 int
}

type region struct {
	Southern, Central, Northern bool
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	Hotels := []hotel{
		hotel{
			Name:    "Hilton",
			Address: "1000 Hotel Street",
			City:    "Los Angeles",
			Region: region{
				Southern: true,
				Central:  false,
				Northern: false,
			},
			Zip: 34567,
		},
		hotel{
			Name:    "Embassy",
			Address: "4200 Suit Drive",
			City:    "San Fransico",
			Region: region{
				Southern: false,
				Central:  false,
				Northern: true,
			},
			Zip: 54321,
		},
	}

	err := tpl.Execute(os.Stdout, Hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
