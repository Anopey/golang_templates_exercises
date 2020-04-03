package main

import "text/template"

const (
	//Southern signifies southern California
	Southern Region = "Southern"
	//Northern signifies Northern California
	Northern Region = "Northern"
	//Central signifies Central California
	Central Region = "Central"
)

//Region defines a region in California
type Region string

type hotel struct {
	Name, Address, City, Zip string
	HotelRegion              Region
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

}
