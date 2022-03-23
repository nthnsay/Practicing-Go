package main

import (
	"encoding/json"
	"fmt"

	//"html/template"
	//"encoding/json"
	"os"
	"flag"
)

func main() {
	filename := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *filename)
	
	f,err := os.Open(*filename)
	if err !=nil{
		// handle instead of panic
		panic(err)
	}

	d := json.NewDecoder(f)
	var story Story
	
	if err := d.Decode(&story); err !=nil{
		panic(err)
	}

	fmt.Print("%+v\n", story)

}

type Story map[string]Chapter

type Chapter struct {
	Title string `json:"title"`
	Paragraphs []string `json:"story"`
	Options []Option `josn:"options"`
}

type Option struct {
	Text string `json:"text"`
	Chapter  string `json:"arc"`
}
