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

	var partTwoFlag bool
	flag.BoolVar(&partTwoFlag, "partTwo", false, "run part two of the puzzle")
	flag.Parse()


	var screen [6][50]bool
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
			rotateColumn(col, step, &screen)

		} else if strings.HasPrefix(line, "rotate r") {
			n := strings.Split(line[13:], " by ")
			row, _ := strconv.Atoi(n[0])
			step, _ := strconv.Atoi(n[1])
			rotateRow(row, step, &screen)
		}
	}

	if partTwoFlag {
		printScreen(&screen)
	} else {
		println(countLitPixels(&screen))
	}

}

func rotateColumn(column int, steps int, screen *[6][50]bool) {
	var temp [6]bool
	for i := 0; i < 6; i++ {
		temp[i] = screen[i][column]
	}
	for index, value := range temp {
		screen[(index + steps) % 6][column] = value
	}
}

func rotateRow(row int, steps int, screen *[6][50]bool) {
	var temp [50]bool
	for i := 0; i < 50; i++ {
		temp[i] = screen[row][i]
	}
	for index, value := range temp {
		screen[row][(index + steps) % 50] = value
	}
}

func printScreen(screen *[6][50]bool) {
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

func turnOnPixels(width int, height int, screen *[6][50]bool) {
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			screen[h][w] = true
		}
	}
}

func countLitPixels(screen *[6][50]bool) (sum int) {
	sum = 0
	for _, row := range screen {
		for _, pixel := range row {
			if pixel {
				sum++
			}
		}
	}
	return sum
}