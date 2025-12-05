package day02

import (
	"aoc2025/utils"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

//go:embed example.txt
var _ string

func Pt1() {
	ranges := strings.Split(input, ",")
	total := 0
	for _, s := range ranges {
		pts := strings.Split(s, "-")
		start, end := utils.MustAtoi(pts[0]), utils.MustAtoi(pts[1])
		if end <= start {
			panic("invalid range")
		}
		for ; start <= end; start++ {
			if isRepeat(strconv.Itoa(start)) {
				total += start
			}
		}
	}
	fmt.Println(total)
}

func isRepeat(s string) bool {
	if len(s) <= 1 {
		return false
	}
	if len(s)%2 != 0 {
		return false
	}

	r := []rune(s)
	mid := len(r) / 2

	left := string(r[:mid])
	right := string(r[mid:])
	return left == right
}

func isRepeatAny(s string) bool {
	if len(s) <= 1 {
		return false
	}

	for size := 1; size < len(s); size++ {
		if len(s)%size != 0 {
			continue
		}
		partNum := len(s) / size
		parts := make([]string, partNum)
		for j := 0; j < partNum; j++ {
			start := j * size
			end := start + size
			parts[j] = s[start:end]
		}
		allEquals := true
		for k := 1; k < len(parts); k++ {
			if parts[k] != parts[0] {
				allEquals = false
				break
			}
		}
		if allEquals {
			return true
		}
	}
	return false
}

func Pt2() {
	ranges := strings.Split(input, ",")
	total := 0
	for _, s := range ranges {
		pts := strings.Split(s, "-")
		start, end := utils.MustAtoi(pts[0]), utils.MustAtoi(pts[1])
		if end <= start {
			panic("invalid range")
		}
		for ; start <= end; start++ {
			if isRepeatAny(strconv.Itoa(start)) {
				//fmt.Printf("found repeat any: %d\n", start)
				total += start
			}
		}
	}
	fmt.Println(total)
}
