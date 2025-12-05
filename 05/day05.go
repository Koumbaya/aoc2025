package day05

import (
	"aoc2025/utils"
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

//go:embed example.txt
var _ string

type rge struct {
	start int
	end   int
}

func Pt1() {
	ids, ingredients := utils.LinesTwoParts(input)
	fresh := make([]rge, len(ids))
	for i, f := range ids {
		start, end := utils.SplitInTwoInts(f, "-")
		fresh[i] = rge{start: start, end: end}
	}
	res := 0
	for _, s := range ingredients {
		num := utils.MustAtoi(s)
		for i := range fresh {
			if num >= fresh[i].start && num <= fresh[i].end {
				res++
				break
			}
		}
	}
	fmt.Println(res)
}

func Pt2() {
	ids, _ := utils.LinesTwoParts(input)
	rg := utils.NewDisjointRange()
	for _, f := range ids {
		start, end := utils.SplitInTwoInts(f, "-")
		rg.AddRange(start, end)
	}

	fmt.Println(rg.SumInclRanges())
}

// 407122360249952 too high
// 393143620099650 too high
