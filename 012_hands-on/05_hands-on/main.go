package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	template.Must(template.ParseFiles("tpl.gohtml"))
}

type dish struct {
	Price float32
	Name  string
}

type menuSection struct {
	SectionName string
	Dishes      []dish
}

func main() {
	var foods = map[string]dish{
		"Eggs": dish{
			Price: 4.99,
			Name:  "Eggs",
		},
		"Toast": dish{
			Price: 7.99,
			Name:  "Toast",
		},
		"Coffee": dish{
			Price: 1.99,
			Name:  "Coffee",
		},
		"Breakfast Plate": dish{
			Price: 9.99,
			Name:  "Breakfast Plate",
		},
		"Hambaga": dish{
			Price: 7.99,
			Name:  "Hambaga",
		},
		"Sushi": dish{
			Price: 11.99,
			Name:  "Sushi",
		},
		"Pizza": dish{
			Price: 7.99,
			Name:  "Pizza",
		},
		"Kebab": dish{
			Price: 12.99,
			Name:  "Kebab",
		},
		"Restuarant": dish{
			Price: 99999.99,
			Name:  "Restuarant",
		},
	}

	var menuSections = []menuSection{
		menuSection{
			SectionName: "Breakfast",
			Dishes: []dish{
				foods["Eggs"],
				foods["Toast"],
				foods["Coffee"],
				foods["Breakfast Plate"],
				foods["Restuarant"],
			},
		},
		menuSection{
			SectionName: "Lunch",
			Dishes: []dish{
				foods["Kebab"],
				foods["Toast"],
				foods["Coffee"],
				foods["Hambaga"],
				foods["Pizza"],
				foods["Restuarant"],
			},
		},
		menuSection{
			SectionName: "Dinner",
			Dishes: []dish{
				foods["Kebab"],
				foods["Sushi"],
				foods["Hambaga"],
				foods["Pizza"],
				foods["Restuarant"],
			},
		},
	}

	err := tpl.Execute(os.Stdout, menuSections)
	if err != nil {
		log.Fatalln(err)
		panic("oh noes!")
	}
}
