package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("Welcome to QuizGame!\n")

	filename := flag.String("file", "problems.csv", "CSV file containing problems in format (question,answer)")
	timeout := flag.Int("timeout", 30, "Timeout period for quiz")

	flag.Parse()

	reader, err := createCSVReader(*filename)

	if err != nil {
		fmt.Printf("Error reading file: %s\n", err.Error())
		os.Exit(1)
	}
	questions := readProblems(reader)
	correctAnswers := askQuestions(questions, *timeout)

	fmt.Printf("\nYour score is: %d/%d\n", correctAnswers, len(questions))
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

func askQuestions(problems [][]string, timeout int) int {
	timer := time.NewTimer(time.Duration(timeout) * time.Second)
	count := 1
	correctAnswers := 0
	answerCh := make(chan string)
problems:
	for _, p := range problems {
		question := fmt.Sprintf("Question: %s", p[0])
		correctAnswer := strings.TrimSpace(p[1])
		fmt.Printf("%d. %s: ", count, question)
		count++
		go func() {
			var answer string
			fmt.Scanf("%s", &answer)
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			fmt.Println("\nTime's up!")
			break problems
		case answer := <-answerCh:
			if correctAnswer == strings.TrimSpace(answer) {
				correctAnswers++
			}
		}
	}
	return correctAnswers
}
