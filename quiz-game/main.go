package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	q string
	a string
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file with questions,answers format")
	timeLimit := flag.Int("limit", 30, "quiz time limit in secons")
	flag.Parse()

	// open and parse cvs problem file
	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open file %s", *csvFileName))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file")
	}
	problems := parseLines(lines)

	// start the timer
	correct := 0
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for i, p := range problems {
		fmt.Printf("Problem #%02d: %s = ", i+1, p.q)
		ansch := make(chan string)
		go func() {
			var ans string
			fmt.Scanf("%s\n", &ans)
			ansch <- ans
		}()

		select {
		case <-timer.C:
			{
				fmt.Printf("\n")
				fmt.Print("Timie finished\n")
				fmt.Printf("Your scored %d out of %d\n", correct, len(problems))
				return
			}
		case ans := <-ansch:
			{
				if ans == p.a {
					correct++
				}
			}
		}
	}

	fmt.Printf("Your scored %d out of %d\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return problems
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
