package day03

import (
	"aoc2025/utils"
	_ "embed"
	"fmt"
	"strconv"
)

//go:embed input.txt
var input string

//go:embed example.txt
var _ string

func Pt1() {
	lines := utils.Lines(input)
	jolts := 0
	for _, line := range lines {
		h1, p1 := utils.HighestNumForward(line[:len(line)-1])
		h2, _ := utils.HighestNumBackward(line[p1+1:])
		//fmt.Println((h1 * 10) + h2)
		jolts += (h1 * 10) + h2
	}
	fmt.Println(jolts)
}

func Pt2() {
	lines := utils.Lines(input)
	jolts := 0
	for _, line := range lines {
		bankStr := ""
		start := 0
		for remaind := 12; remaind > 0; remaind-- { // start from 12 then decrease, that's how many digits we must reserve for the rest of the number
			v := 0
			v, start = findHighestWithinRange(line, start, remaind)
			start++ // move to next position for next search
			bankStr += strconv.Itoa(v)
		}
		bank := utils.MustAtoi(bankStr)
		//fmt.Println(bankStr)
		jolts += bank
	}
	fmt.Println(jolts)
}

func findHighestWithinRange(s string, start, cutoff int) (val int, pos int) {
	val = -1
	for ; start <= len(s)-cutoff; start++ {
		n := utils.MustAtoi(string(s[start]))
		if utils.MustAtoi(string(s[start])) > val {
			val = n
			pos = start
		}
	}
	return val, pos
}
