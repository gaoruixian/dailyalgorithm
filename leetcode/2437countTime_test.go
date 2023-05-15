package leetcode

import (
	"fmt"
	"strconv"
	"testing"
)

func countTime(time string) int {
	a := 24
	b := 1
	c := 1
	bytes := []byte(time)
	if bytes[0] != '?' && bytes[1] != '?' {
		a = 1
	}
	if bytes[0] == '?' && bytes[1] != '?' {
		change, _ := strconv.Atoi(string(bytes[1]))
		if change >= 4 {
			a = 2
		} else {
			a = 3
		}
	}
	if bytes[0] != '?' && bytes[1] == '?' {
		if (bytes[0]) == '2' {
			a = 4
		} else {
			a = 10
		}
	}
	if time[3] == '?' {
		b = 6
	}
	if time[4] == '?' {
		c = 10
	}
	return a * b * c
}

func countTime2(time string) int {
	h := 24
	m := 60
	bytes := []byte(time)
	if bytes[0] != '?' && bytes[1] != '?' {
		h = 1
	}
	if bytes[0] == '?' && bytes[1] != '?' {
		change, _ := strconv.Atoi(string(bytes[1]))
		if change >= 4 {
			h = 2
		} else {
			h = 3
		}
	}
	if bytes[0] != '?' && bytes[1] == '?' {
		if (bytes[0]) == '2' {
			h = 4
		} else {
			h = 10
		}
	}
	if time[3] != '?' {
		m = m / 6
	}
	if time[4] != '?' {
		m = m / 10
	}
	return h * m
}
func TestCount(t *testing.T) {
	time := "?4:21"
	fmt.Println(countTime(time))
	fmt.Println(countTime2(time))
}
func TestByte(t *testing.T) {
	char := '5'
	a, _ := strconv.Atoi(string(char))
	if int(char) > 5 {

		fmt.Println("big:", a, string(char))
	}
}
