package main

import (
	"flag"
	"os"
	"log"
	"io/ioutil"
	"strings"
)

var input string

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	byteArray, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
		return
	}
	input = string(byteArray)

	var partTwoFlag bool
	flag.BoolVar(&partTwoFlag, "partTwo", false, "run part two of the puzzle")
	flag.Parse()

	if partTwoFlag {
		partTwo()
	} else {
		partOne()
	}

}

func partOne() {
	tlsCounter := 0
	for _, row := range strings.Split(input, "\n") {
		outsideBracketABBA := false
		insideBracketABBA := false
		temp := ""
		for index, character := range row {
			if character == '[' || index == len(row) - 1 {
				if containsABBA(temp) {
					outsideBracketABBA = true
				}
				temp = ""
			} else if character == ']' {
				if containsABBA(temp) {
					insideBracketABBA = true
					break
				}
				temp = ""
			} else {
				temp += string(character)
			}
		}
		if outsideBracketABBA && !insideBracketABBA {
			tlsCounter++
		}

	}
	println(tlsCounter)
}
func partTwo() {
	println("partTwo not implemented yet")
}

func containsABBA(input string) bool {
	if len(input) < 4 {
		return false
	}

	for i := 0; i < len(input) - 3; i++ {
		search := input[i] == input[i + 3] && input[i + 1] == input[i + 2] &&
			input[i] != input[i + 2]
		if search {
			return true
		}
	}
	return false
}