package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to QuizGame!\n")

	filename := flag.String("file", "problems.csv", "CSV file containing problems")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err.Error())
		os.Exit(1)
	}

	reader := csv.NewReader(file)

	correctAnswers := 0
	count := 1

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		question := fmt.Sprintf("Question: %s", record[0])
		correctAnswer := record[1]
		var userAnswer string
		fmt.Printf("%d. %s: ", count, question)
		count++
		fmt.Scanf("%s", &userAnswer)
		if correctAnswer == userAnswer {
			correctAnswers++
		}
	}

	fmt.Printf("Your score is: %d/%d\n", correctAnswers, count)
}
