package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/srinchow/quiz/problem"
	"os"
	"time"
)

func parseLines(lines [][]string) []problem.Problem {
	ret := make([]problem.Problem, len(lines))
	for i, line := range lines {
		ret[i] = problem.New(line)
	}
	return ret
}

func main() {

	csvFileName := flag.String("csv", "problem.csv", "Opens the given csv file for question and answers")
	flag.Parse()

	timeLimit := flag.Int("timelimit", 10, "timelimit for the quiz")

	csvFile, err := os.Open(*csvFileName)

	if err != nil {
		exit(fmt.Sprintf("Failed to open the file %v", *csvFile))
	}

	csvReader := csv.NewReader(csvFile)

	lines, err := csvReader.ReadAll()

	if err != nil {
		exit("Error parsing csv file")
	}

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	answerChan := make(chan string)
	problems := parseLines(lines)
	counter := 0

	for i, p := range problems {
		fmt.Printf("Problem %v, %v = ", i+1, p.Question)
		go waitForAnswer(answerChan)
		select {
		case <-timer.C:
			fmt.Printf("\nTotal score %v out of %v questions \n", counter, len(problems))
			return
		case answer := <-answerChan:
			if answer == p.Answer {
				counter++
			}
		}
	}
	fmt.Printf("Total score %v out of %v questions \n", counter, len(problems))
}

func waitForAnswer(ch chan string) {
	var answer string
	fmt.Scanf("%s\n", &answer)
	ch <- answer
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
