package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type Seconds int

func main() {
	//Take in input flag to set timer
	seconds := flag.Int("seconds", 30, "Set amount of time for quiz")
	flag.Parse()

	//Get user input here to start timer
	fmt.Print("Press enter to start: ")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	//Timer starts

	csv := csvReader()
	timer := time.NewTicker(time.Duration(*seconds) * time.Second)

	var counter int
	for _, csvSlice := range csv {
		fmt.Println(csvSlice[0])

		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Print("Enter your answer:")
			fmt.Scanln(&answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println("Times up!")
			fmt.Printf("You got %d out of %d questions correct", counter, len(csv))
			return
		case answer := <-answerCh:
			if answer == csvSlice[1] {
				counter++
			}
		}

	}
	fmt.Printf("You got %d out of %d questions correct", counter, len(csv))
}
func csvReader() [][]string {
	//open file
	recordFile, err := os.Open("./problems.csv")
	if err != nil {
		fmt.Println("An error happened ::", err)
	}
	//Initialise the reader
	reader := csv.NewReader(recordFile)
	// reads all the records
	records, _ := reader.ReadAll()
	// Iterate trhough the records
	return records
}
