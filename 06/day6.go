package main

import (
	"flag"
	"os"
	"log"
	"io/ioutil"
	"strings"
)

var input string
var errorCorrectedMessage string

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
	numberOfLines := strings.Count(input, "\n")
	numberOfCharactersPerRow := strings.Index(input, "\n")
	var mostCommonCharacter rune
	var mostCommonCharacterCount int

	for position := 0; position < numberOfCharactersPerRow; position++ {
		mostCommonCharacterCount = 0
		for character := 'a'; character <= 'z'; character++ {
			characterCount := 0
			for row := 0; row < numberOfLines; row++ {
				if []rune(input)[position + (numberOfCharactersPerRow + 1) * row] == character {
					characterCount++
				}
			}
			if characterCount > mostCommonCharacterCount {
				mostCommonCharacterCount = characterCount
				mostCommonCharacter = character
			}
		}
		errorCorrectedMessage += string(mostCommonCharacter)
	}
	println(errorCorrectedMessage)
}
func partTwo() {
	numberOfLines := strings.Count(input, "\n")
	numberOfCharactersPerRow := strings.Index(input, "\n")
	var leastCommonCharacter rune
	var leastCommonCharacterCount int

	for position := 0; position < numberOfCharactersPerRow; position++ {
		leastCommonCharacterCount = numberOfLines + 1
		for character := 'a'; character <= 'z'; character++ {
			characterCount := 0
			for row := 0; row < numberOfLines; row++ {
				if []rune(input)[position + (numberOfCharactersPerRow + 1) * row] == character {
					characterCount++
				}
			}
			if characterCount < leastCommonCharacterCount {
				leastCommonCharacterCount = characterCount
				leastCommonCharacter = character
			}
		}
		errorCorrectedMessage += string(leastCommonCharacter)
	}
	println(errorCorrectedMessage)
}
