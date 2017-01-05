package main

import (
	"flag"
	"os"
	"log"
	"io/ioutil"
	"strings"
	"strconv"
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
	input = "rect 3x2\nrotate column x=1 by 49"

	var partTwoFlag bool
	flag.BoolVar(&partTwoFlag, "partTwo", false, "run part two of the puzzle")
	flag.Parse()

	if partTwoFlag {
		partTwo()
	} else {
		partOne()
	}

}

//rect 1x1
//rotate row y=0 by 10
//rect 1x1
//rotate row y=0 by 10
//rect 1x1
//rotate row y=0 by 5
//rect 1x1
//rotate row y=0 by 3
//rect 2x1
//rotate row y=0 by 4
//rotate column x=1 by 1
//rotate row y=0 by 4

func partOne() {
	var screen [50][6]bool
	for _, line := range strings.Split(input, "\n") {
		if strings.HasPrefix(line, "rect") {
			n := strings.Split(line[5:], "x")
			width, _ := strconv.Atoi(n[0])
			height, _ := strconv.Atoi(n[1])
			turnOnPixels(width, height, &screen)
		} else if strings.HasPrefix(line, "rotate c") {
			n := strings.Split(line[16:], " by ")
			col, _ := strconv.Atoi(n[0])
			step, _ := strconv.Atoi(n[1])
			printScreen(&screen)
			rotateColumn(col, step, &screen)
			printScreen(&screen)

		} else if strings.HasPrefix(line, "rotate r") {

		}
	}

}
func rotateColumn(column int, steps int, screen *[50][6]bool) {
	var temp [50]bool
	for i := 0; i < 50; i++{
		temp[i] = screen[i][column]
	}
	for index, value := range temp {
		screen[(index + steps) % 50][column] = value
	}
}
func partTwo() {
	println("partTwo not implemented yet")
}

func printScreen(screen *[50][6]bool) {
	for w := 0; w < len(screen); w++ {
		for h := 0; h < len(screen[0]); h++ {
			if screen[w][h] {
				print("#")
			} else {
				print(".")
			}
		}
		println("")
	}

	println("-----------")
}

func turnOnPixels(width int, height int, screen *[50][6]bool) {
	for w := 0; w < width; w++ {
		for h := 0; h < height; h++ {
			screen[h][w] = true
		}
	}
}