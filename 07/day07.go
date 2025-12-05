package day07

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
	s := g.Find("S")
	dirs := utils.GetCardinals()
	split := walk(s, g, dirs)
	fmt.Println(split)
}

func walk(p utils.Pos, g *utils.Grid, dirs map[utils.DirName]utils.Direction) (split int) {
	if !g.Valid(p) {
		return 0
	}
	s := g.Get(p)
	switch s {
	case "S":
		//start, go south
		return walk(g.GetPosFromDirs(p, []utils.Direction{dirs[utils.S]})[0], g, dirs)
	case "^":
		// try left and right
		split++
		split += walk(g.GetPosFromDirs(p, []utils.Direction{dirs[utils.E]})[0], g, dirs)
		split += walk(g.GetPosFromDirs(p, []utils.Direction{dirs[utils.W]})[0], g, dirs)
		return split
	case ".":
		// continue south
		g.Set(p, "|")
		return walk(g.GetPosFromDirs(p, []utils.Direction{dirs[utils.S]})[0], g, dirs)
	case "|":
		// already visited
		return 0
	default:
		return 0
	}
}

func walk2(p utils.Pos, g *utils.Grid, dirs map[utils.DirName]utils.Direction) (split int) {
	if v := g.CountVisitedTimes(utils.PosDir{Pos: p}); v > 0 {
		return v
	}
	if !g.Valid(p) {
		return 1 // bottom
	}
	g.Visit(utils.PosDir{Pos: p})

	s := g.Get(p)
	switch s {
	case "^":
		// try left and right
		split += walk2(g.GetPosFromDirs(p, []utils.Direction{dirs[utils.E]})[0], g, dirs)
		split += walk2(g.GetPosFromDirs(p, []utils.Direction{dirs[utils.W]})[0], g, dirs)
	default:
		// continue south
		split += walk2(g.GetPosFromDirs(p, []utils.Direction{dirs[utils.S]})[0], g, dirs)
	}
	g.SetVisitedCount(utils.PosDir{Pos: p}, split)
	return split
}

func Pt2() {
	g := utils.NewGrid(utils.Lines(input))
	s := g.Find("S")
	dirs := utils.GetCardinals()
	split := walk2(s, g, dirs)
	fmt.Println(split)
}

// 3270 too low
