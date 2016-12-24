package main

import (
	"os"
	"log"
	"io/ioutil"
	"fmt"
	"strings"
	"strconv"
	"github.com/bradfitz/slice"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	byte, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
		return
	}
	input := string(byte)
	numberOfOccrencies := make(map[string]int)

	for _, row := range strings.Split(input, "\n") {
		numberOfOccrencies := calculateNumberOfOccurrences(row)
		for index, nbr := range numberOfOccrencies {
			print(string(index + 61) + " -> " + strconv.Itoa(nbr))
		}
	}
	for key := range numberOfOccrencies {
		fmt.Printf("%s\n", string(key))

	}
}

func calculateNumberOfOccurrences(string string) [28]int {
	//alphabet := [...]string{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	numberOfOccurrences := [28]int{0}
	for index, character := range string {
		if character == '['{
			break;
		}
		if character < 'a'{
			continue
		}
		fmt.Printf("%c, %v, %v\n", string[index], character, character - 'a')
		numberOfOccurrences[character - 'a']++
	}

	return numberOfOccurrences
}

func getChecksum(numberOfOccurences []int) string{
	alphabet := [...]string{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

	//get top five characters
	top5Indexes := make(map[string]int)
	for index, _ := range numberOfOccurences{

	}

	return ""
}