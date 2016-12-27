package haggan

import (
	"crypto/md5"
	"fmt"
	"strings"
	"testing"
)

var input = "reyedfim"
var counter int
var password string

func BenchmarkNoStringConvert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		counter = 0
		password = ""
		for {
			md5array := md5.Sum([]byte(fmt.Sprintf("%s%d", input, counter)))
			if md5array[0] == 0 && md5array[1] == 0 && md5array[2] < 16 {
				password += fmt.Sprintf("%x", md5array[2])
			}
			if len(password) > 7 {
				break
			}
			counter++
		}
	}
}

func BenchmarkStringConvert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		counter = 0
		password = ""
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
	}
}
