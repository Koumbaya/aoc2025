package day04

import (
	"aoc2025/utils"
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

//go:embed example.txt
var _ string

func Pt1() {
	g := utils.NewGrid(utils.Lines(input))
	ats := g.FindAll("@")
	total := 0
	for _, pos := range ats {
		n := g.CountNeighborsIn(pos, "@")
		if n < 4 {
			total++
		}
	}
	fmt.Println(total)
}

func Pt2() {
	g := utils.NewGrid(utils.Lines(input))
	total := 0
	for {
		totalLoop := 0
		ats := g.FindAll("@")
		for _, pos := range ats {
			n := g.CountNeighborsIn(pos, "@")
			if n < 4 {
				total++
				totalLoop++
				g.Set(pos, "x")
			}
		}
		if totalLoop == 0 {
			break
		}
	}
	fmt.Println(total)

}
