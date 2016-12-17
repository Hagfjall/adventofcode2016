package main

import (
	"io/ioutil"
	"os"
	"log"
	"strings"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	input, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
		return
	}
	s := string(input)
	numberOfTraingles := 0
	for _, instruction := range strings.Split(s, "\n") {
		if isTriangle(strings.Fields(instruction)) {
			numberOfTraingles++
		}
	}
	print(numberOfTraingles)
}

func isTriangle(fields []string) bool {
	var triangle [3]int
	for index, field := range fields {
		triangle[index], _ = strconv.Atoi(field)
	}
	return triangle[0] + triangle[1] > triangle[2] &&
		triangle[1] + triangle[2] > triangle[0] &&
		triangle[0] + triangle[2] > triangle[1]
}