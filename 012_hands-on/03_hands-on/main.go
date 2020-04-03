package main

import (
	"log"
	"os"
	"text/template"
)

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
	hotels := []hotel{
		hotel{
			Name:        "Hotel California",
			Address:     "42 Sunset Boulevard",
			City:        "Los Angeles",
			Zip:         "95612",
			HotelRegion: Southern,
		},
		hotel{
			Name:        "H",
			Address:     "4",
			City:        "L",
			Zip:         "95612",
			HotelRegion: Southern,
		},
	}

	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
