package main

import (
	"strings"
	"fmt"
	"strconv"
	"math"
)

func main() {
	input := "L2, L3, L3, L4, R1, R2, L3, R3, R3, L1, L3, R2, R3, L3, R4, R3, R3, L1, L4, R4, L2, R5, R1, L5, R1, R3, L5, R2, L2, R2, R1, L1, L3, L3, R4, R5, R4, L1, L189, L2, R2, L5, R5, R45, L3, R4, R77, L1, R1, R194, R2, L5, L3, L2, L1, R5, L3, L3, L5, L5, L5, R2, L1, L2, L3, R2, R5, R4, L2, R3, R5, L2, L2, R3, L3, L2, L1, L3, R5, R4, R3, R2, L1, R2, L5, R4, L5, L4, R4, L2, R5, L3, L2, R4, L1, L2, R2, R3, L2, L5, R1, R1, R3, R4, R1, R2, R4, R5, L3, L5, L3, L3, R5, R4, R1, L3, R1, L3, R3, R3, R3, L1, R3, R4, L5, L3, L1, L5, L4, R4, R1, L4, R3, R3, R5, R4, R3, R3, L1, L2, R1, L4, L4, L3, L4, L3, L5, R2, R4, L2"
	direction := 0 //start pointing north
	north := 0
	east := 0

	for _, instruction := range strings.Split(input, ", ") {
		if strings.HasPrefix(instruction, "R") {
			direction += 1
			direction %= 4
		} else if strings.HasPrefix(instruction, "L") {
			direction -= 1
			if direction < 0{
				direction = 3
			}
		}
		direction = int(math.Abs(float64(direction) ))
		steps, _ := strconv.Atoi(instruction[1:])
		if direction == 0 {
			//north
			north += steps
		} else if direction == 1 {
			//east
			east += steps
		} else if direction == 2 {
			//south
			north -= steps
		} else if direction == 3 {
			//west
			east -= steps
		}
	}
	fmt.Printf("%d blocks from start", north + east)


}