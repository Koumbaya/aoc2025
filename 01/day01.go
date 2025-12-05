package day01

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
	in := utils.Lines(input)
	dial := 50
	zeroes := 0
	for _, l := range in {
		dir := l[0]
		num := utils.MustAtoi(l[1:])
		num %= 100
		switch dir {
		case 'L':
			dial -= num
		case 'R':
			dial += num
		default:
			panic("unknown direction")
		}
		if dial > 99 {
			dial -= 100
		} else if dial < 0 {
			dial += 100
		}
		if dial == 0 {
			zeroes++
		}
	}
	fmt.Println(zeroes)
}

func Pt2() {
	in := utils.Lines(input)
	dial := 50
	prevDial := dial
	zeroes := 0
	for _, l := range in {
		dir := l[0]
		fnum := utils.MustAtoi(l[1:])
		rem := fnum / 100
		num := fnum % 100
		switch dir {
		case 'L':
			dial -= num
		case 'R':
			dial += num
		default:
			panic("unknown direction")
		}
		if dial > 99 {
			dial -= 100
			if prevDial != 0 {
				zeroes++
			}
		} else if dial < 0 {
			dial += 100
			if prevDial != 0 {
				zeroes++
			}
		} else if dial == 0 {
			zeroes++
		}
		zeroes += rem

		fmt.Printf("rem: %d, num: %d, dial: %d,   zeroes: %d\n", rem, num, dial, zeroes)
		prevDial = dial
	}
	fmt.Println(zeroes)
}
