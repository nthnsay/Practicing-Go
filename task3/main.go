package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	//"html/template"
	//"encoding/json"
	"os"
)

func main() {
	var storyContainer story
	jsonFile, err := os.Open("./gopher.json")

	if err != nil {
		fmt.Println(err)
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}
	//var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &storyContainer)
	fmt.Println(storyContainer)
	//err2 := json.Unmarshal(jsonFile, &stories)

	//Read in Json file into a map of string[storyarc]
}

type story struct {
	storyarcs intro
}

type intro struct {
	title string
	story []string
}

type options struct {
	text string
	arc  string
}
