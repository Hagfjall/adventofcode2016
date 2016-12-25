package main

import (
	"os"
	"log"
	"io/ioutil"
	"fmt"
	"strings"
	"sort"
	"github.com/cznic/sortutil"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
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
	sectorIdSum := 0

	for _, row := range strings.Split(input, "\n") {
		calculatedChecksum := calculateChecksum(row)
		checksum := row[strings.LastIndex(row, "[") + 1:len(row) - 1]
		if checksum == calculatedChecksum {
			sectorId, _ := strconv.Atoi(row[strings.LastIndex(row, "-") + 1:strings.LastIndex(row, "[")])
			sectorIdSum += sectorId
		}
	}
	print(sectorIdSum)
}

func calculateChecksum(string string) string {
	numberOfOccurrences := make(map[rune]int)

	//count number of occurrences for all characters
	for _, character := range string {
		if character == '[' { //read all characters before checksum
			break;
		}
		if character < 'a' { //counts for '-' as well as all numerics
			continue
		}
		numberOfOccurrences[character]++
	}

	//sort by occurrences and alphabetically
	checksum := ""
	n := map[int][]rune{}
	var a []int
	for k, v := range numberOfOccurrences {
		n[v] = append(n[v], k)
	}
	for k := range n {
		a = append(a, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	c := 0
	Loop:
	for _, k := range a {
		sort.Sort(sortutil.RuneSlice(n[k]))
		for _, s := range n[k] {
			if c > 4 {
				break Loop
			}

			checksum += fmt.Sprintf("%c", s)
			c++
		}
	}

	return checksum
}