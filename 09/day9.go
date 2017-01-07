package main

import (
	"flag"
	"os"
	"log"
	"io/ioutil"
	"strings"
	"strconv"
	"bytes"
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
	input = "X(8x2)(3x3)ABCY"

	var partTwoFlag bool
	flag.BoolVar(&partTwoFlag, "partTwo", false, "run part two of the puzzle")
	flag.Parse()

	if partTwoFlag {
		partTwo()
	} else {
		partOne()
	}

}
//examples
//A(1x5)BC = ABBBBBC
//(3x3)XYZ = XYZXYZXYZ
//A(2x2)BCD(2x2)EFG = ABCBCDEFEFG
//X(8x2)(3x3)ABCY = X(3x3)ABC(3x3)ABCY
func partOne() {
	var decompressedData bytes.Buffer
	for index := 0; index < len(input); index++ {
		if input[index] == '(' {
			nextEndParenthesis := strings.Index(input[index + 1:], ")") + index + 1
			n := strings.Split(input[index + 1:nextEndParenthesis], "x")
			nbrOfCharacters, _ := strconv.Atoi(n[0])
			nbrOfTimes, _ := strconv.Atoi(n[1])
			for nbrOfTimes > 0 {
				decompressedData.WriteString(input[nextEndParenthesis + 1: nextEndParenthesis + nbrOfCharacters + 1])
				nbrOfTimes--
			}
			index = nextEndParenthesis + nbrOfCharacters;
		} else {
			decompressedData.WriteByte(input[index])
		}

	}
	println(decompressedData.String())
}
func partTwo() {
	println("partTwo not implemented yet")
}
