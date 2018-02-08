package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type quiz struct {
	Question string `json:"question,omitempty"`
	Answer   string `json:"answer,omitempty"`
}

var quizz []quiz
var correct int
var wrong int

func main() {
	file, err := os.Open("problems.csv")
	reader := csv.NewReader(bufio.NewReader(file))

	if err != nil {
		os.Exit(1)
	}

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		quizz = append(quizz, quiz{
			Question: line[0],
			Answer:   line[1],
		},
		)
	}

	for i := range quizz {
		fmt.Println(quizz[i].Question)
		fmt.Println(quizz[i].Answer)

		for {
			argsWithoutProg := os.Args[1:]
			if argsWithoutProg[i] == quizz[i].Answer {
				correct++
			} else {
				wrong++
			}
		}
	}

}
