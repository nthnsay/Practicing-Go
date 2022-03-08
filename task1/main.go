package main

import (
	"encoding/csv"
	"fmt"
	"os"
)
  func main() {
	csv:=csvReader()
	var counter int	
	for _, csvSlice := range csv{
		fmt.Println( csvSlice[0])
		fmt.Print("Enter your answer:")
		var answer string
		fmt.Scanln(&answer)
		if answer == csvSlice[1]{
			counter ++
		}
		
	}
	fmt.Printf("You got %d out of %d questions correct", counter, len(csv))
  }
  func csvReader() [][]string  {
		//open file
		recordFile, err := os.Open("./problems.csv")
		if err !=nil{
			fmt.Println("An error happened ::", err)
		}
		//Initialise the reader
		reader := csv.NewReader(recordFile)
		// reads all the records	
		records, _ := reader.ReadAll()
		// Iterate trhough the records 
		return records
}