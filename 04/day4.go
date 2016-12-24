package main

import (
	"os"
	"log"
	"io/ioutil"
	"fmt"
	"strings"
	"sort"
	"github.com/cznic/sortutil"
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
	//numberOfOccrencies := make(map[string]int)

	for _, row := range strings.Split(input, "\n") {
		print(calculateChecksum(row))
		//for index, nbr := range numberOfOccrencies {
		//	//print(string(index + 61) + " -> " + strconv.Itoa(nbr))
		//}
	}
	//for key := range numberOfOccrencies {
	//	fmt.Printf("%s\n", string(key))
	//
	//}
}

//aaaaa-bbb-z-y-x-123[abxyz]
func calculateChecksum(string string) string {
	//alphabet := [...]string{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	numberOfOccurrences := make(map[rune]int)
	//count number of occurrences for all characters
	for _, character := range string {
		if character == '[' {
			break;
		}
		if character < 'a' {
			continue
		}
		//fmt.Printf("%c, %v, %v\n", string[index], character, character - 'a')
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
			fmt.Printf("%d: %c, %d\n", c, s, k)
			c++
		}
	}

	return checksum
}