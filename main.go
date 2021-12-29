package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

var (
	filename  = flag.String("filename", "problems.csv", "path to CSV file containing the quiz")
	timeout   = flag.Int("timeout", 30, "time in seconds to wait for every question")
	questions = 0
	correct   = 0
	wrong     = 0
)

func main() {
	flag.Parse()

    if *timeout < 1 {
        err := fmt.Sprintf("-timeout flag must be set to a value greater than 0. Given was %d\n", *timeout)
        panic(err)
    }

	fmt.Printf("Opening %s\n", *filename)
	file, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	r := csv.NewReader(file)
	input := make(chan string, 1)

	for {
		record, err := r.Read()
		questions++
		if err != nil {
			if err == io.EOF {
				err = nil
				break
			}
		}

		question := record[0]
		answer := record[1]
		fmt.Printf("Question %d:\n%s\n", questions, question)
		fmt.Printf("Your answer: ")
		go getInput(input)
		var resp string
		select {
		case resp = <-input:
			fmt.Println(resp)
		case <-time.After(time.Duration(*timeout) * 1000 * time.Millisecond):
			fmt.Printf("Timeout!\n")

		}

		if resp == answer {
			fmt.Println("Correct!\n")
			correct++
		} else {
			fmt.Printf("Wrong. Answer is [%s]. Your input was [%s]\n", answer, resp)
			wrong++
		}

	}
	fmt.Printf("===Results===\n")
	fmt.Printf("Correct: %d\n", correct)
	fmt.Printf("Wrong: %d\n", wrong)
}

func getInput(ch chan string) {
	var input string
	fmt.Scanln(&input)
	ch <- input
}
