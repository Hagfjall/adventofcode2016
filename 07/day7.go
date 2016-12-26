package main

import (
	"flag"
	"os"
	"log"
	"io/ioutil"
	"strings"
	"fmt"
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
		fromIndex := 0
		for fromIndex != -1 {
			println(fmt.Sprintf("fromIndex %d len %d", fromIndex, len(row[fromIndex:])))
			index := strings.Index(row[fromIndex:], "[")
			print(index)
			println(" " + row[fromIndex:])
			checkHypernet := false
			if containsABBA(row[fromIndex:index]) {
				checkHypernet = true
				println("true")
			}else{
				println("false")
			}
			if checkHypernet {
				//index2 := strings.Index(row[fromIndex:], "]")
				//hyperNextContainsAbba := false
				//if containsABBA(row[fromIndex:index2]) {
				//	hyperNextContainsAbba = true
				//}
				//if !hyperNextContainsAbba {
				//	println(row)
				//	tlsCounter++
				//} else {
				//	println("\t" + row)
				//}
			}
			temp := strings.Index(row[fromIndex:], "[")
			println(fmt.Sprintf("temp %d", temp+fromIndex))
			if temp == -1{
				break
			}
			fromIndex += temp + 1
		}

	}
	println(tlsCounter)
}
func partTwo() {
	println("partTwo not implemented yet")
}

func containsABBA(input string) bool {
	println("containsABBA input: " + input)
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