package main

import (
	"strings"
	"crypto/md5"
	"fmt"
	"flag"
)

var input = "reyedfim"
var counter int
var password string

func main() {

	var extra bool
	flag.BoolVar(&extra, "extra", false, "run extra (part two) of the puzzle")
	flag.Parse()
	var f func()
	if extra {
		f = partTwo()
	} else {
		f = partOne()
	}
	f()

}

func partOne() (f func()) {
	f = func() {
		for {
			md5array := md5.Sum([]byte(fmt.Sprintf("%s%d", input, counter)))
			md5 := fmt.Sprintf("%x", md5array)
			if strings.HasPrefix(md5[:], "00000") {
				password += fmt.Sprintf("%c", md5[5])
			}
			if len(password) > 7 {
				break
			}
			counter++
		}
		println(password)

	}
	return f
}

func partTwo() (f func()) {
	foundCharacters := 0
	password = "********"
	f = func() {
		for {
			md5array := md5.Sum([]byte(fmt.Sprintf("%s%d", input, counter)))
			md5 := fmt.Sprintf("%x", md5array)
			if strings.HasPrefix(md5[:], "00000") {
				if '0' <= md5[5] && md5[5] <= '7' {
					index := (md5[5] - '0')
					if password[index] == '*' {
						password = password[:index] + string(md5[6]) + password[index + 1:]
						foundCharacters++
					}
				}
			}
			if foundCharacters > 7 {
				break
			}
			counter++
		}
		println(password)
	}
	return f
}