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
	file, err := os.Open("test.txt")
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
		index := strings.Index(row, "[")
		checkHypernet := false
		for i := 0; i < index; i++ {
			if containsABBA(row[i:index]) {
				checkHypernet = true
				break
			}
		}
		if checkHypernet {
			index2 := strings.LastIndex(row, "]")
			hyperNextContainsAbba := false
			for i := index + 1; i < index2; i++ {
				if containsABBA(row[i:index2]) {
					hyperNextContainsAbba = true
				}
			}
			if !hyperNextContainsAbba {
				println(row)
				tlsCounter++
			} else {
				println("\t" + row)
			}
		}
	}
	println(tlsCounter)
}
func partTwo() {
	println("partTwo not implemented yet")
}

func containsABBA(input string) bool {
	println(input)
	if len(input) < 4 {
		return false
	}
	for i := 0; i < len(input) - 3; i++ {
		search := input[i] == input[i + 3] && input[i + 1] == input[i + 2] &&
			input[i] != input[i + 2]
		if search{
			return true
		}
	}
	return false
}