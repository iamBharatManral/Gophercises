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

	reader, err := createCSVReader(*filename)

	if err != nil {
		fmt.Printf("Error reading file: %s\n", err.Error())
		os.Exit(1)
	}

	totalQuestions, correctAnswers := askQuestions(readProblems(reader))

	fmt.Printf("\nYour score is: %d/%d\n", correctAnswers, totalQuestions)
}

func createCSVReader(filename string) (*csv.Reader, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("Error reading file: %s\n", err.Error())
	}

	return csv.NewReader(file), nil
}

func readProblems(r *csv.Reader) [][]string {
	var problems [][]string
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		problems = append(problems, record)
	}
	return problems
}

func askQuestions(problems [][]string) (int, int) {
	count := 1
	correctAnswers := 0
	for _, p := range problems {
		question := fmt.Sprintf("Question: %s", p[0])
		correctAnswer := p[1]
		var userAnswer string
		fmt.Printf("%d. %s: ", count, question)
		count++
		fmt.Scanf("%s", &userAnswer)
		if correctAnswer == userAnswer {
			correctAnswers++
		}
	}
	return count, correctAnswers
}
