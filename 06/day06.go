package day06

import (
	"aoc2025/utils"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

//go:embed example.txt
var _ string

func Pt1() {
	lines := utils.Lines(input)
	nbDataLines := len(lines) - 1 // exclude operation line
	nums := make([][]int, nbDataLines)
	for i := 0; i < nbDataLines; i++ {
		s := strings.TrimSpace(lines[i])
		nums[i] = utils.SplitToInts(utils.ConcatRunes(s, " ", 4), " ")
	}
	s := strings.TrimSpace(lines[nbDataLines])
	s = utils.ConcatRunes(s, " ", 6)
	ops := strings.Split(s, " ")
	total := 0
	for i, op := range ops {
		switch op {
		case "+":
			t := 0
			for j := 0; j < nbDataLines; j++ {
				t += nums[j][i]
			}
			total += t
		case "*":
			t := 1
			for j := 0; j < nbDataLines; j++ {
				x := nums[j][i]
				t *= x
			}
			total += t
		}
	}
	fmt.Println(total)
}

func Pt2() {
	lines := utils.Lines(input)
	nbCols := len(lines[0]) - 1
	nbLines := len(lines)
	emptyCol := strings.Repeat(" ", nbLines)
	total := 0
	numbers := make([]int, 0)
	// iterate over columns from right to left
	for c := nbCols; c >= 0; c-- {
		// iterate over lines, top to bottom
		s := "" // construct string
		for l := 0; l < nbLines; l++ {
			s += string(lines[l][c])
		}
		if s == emptyCol {
			// separator line, reset numbers
			numbers = make([]int, 0)
			continue
		}
		snum := strings.ReplaceAll(s[0:len(s)-1], " ", "")
		numbers = append(numbers, utils.MustAtoi(snum))

		switch s[nbLines-1] {
		case '+':
			t := 0
			for _, number := range numbers {
				t += number
			}
			total += t
		case '*':
			t := 1
			for _, number := range numbers {
				t *= number
			}
			total += t
		}
	}
	fmt.Println(total)
}
