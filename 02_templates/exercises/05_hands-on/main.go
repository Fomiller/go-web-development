package main

import (
	"html/template"
	"log"
	"os"
)

type item struct {
	Name  string
	Price float64
}

type meal struct {
	Meal  string
	Items []item
}

type menu []meal

type Resturants []resturant

type resturant struct {
	Name string
	Menu menu
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	r1 := Resturants{
		resturant{
			Name: "Chilis",
			Menu: menu{
				meal{
					Meal: "Breakfast",
					Items: []item{
						item{
							Name:  "Eggs",
							Price: 1.99,
						},
						item{
							Name:  "Steak and eggs",
							Price: 6.99,
						},
					},
				},
				meal{
					Meal: "Lunch",
					Items: []item{
						item{
							"Burger",
							12.99,
						},
						item{
							"Buritto",
							8.99,
						},
					},
				},
				meal{
					Meal: "Dinner",
					Items: []item{
						item{
							"Spaghetti and Meatballs",
							14.99,
						},
						item{
							"Risotto",
							21.99,
						},
					},
				},
			},
		},
		resturant{
			Name: "Bobs Burgers",
			Menu: menu{
				meal{
					Meal: "Breakfast",
					Items: []item{
						item{
							Name:  "Grits",
							Price: 3.99,
						},
						item{
							Name:  "Eggs",
							Price: 6.99,
						},
					},
				},
				meal{
					Meal: "Lunch",
					Items: []item{
						item{
							"BOB Burger",
							12.99,
						},
						item{
							"Tacos",
							8.99,
						},
					},
				},
				meal{
					Meal: "Dinner",
					Items: []item{
						item{
							"Steak",
							24.99,
						},
						item{
							"Lobster",
							32.99,
						},
					},
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, r1)
	if err != nil {
		log.Fatalln(err)
	}
}
