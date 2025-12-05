package utils

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func MustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func MustAois(s []string) []int {
	r := make([]int, len(s))
	for i, s := range s {
		r[i] = MustAtoi(s)
	}
	return r
}

func Abs(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

// IntConcat a=10, b=3 returns 103.
func IntConcat(a, b int) int {
	return MustAtoi(fmt.Sprintf("%d%d", a, b))
}

func AllIncreasing[T cmp.Ordered](in []T) bool {
	for i := 1; i < len(in); i++ {
		if in[i] < in[i-1] {
			return false
		}
	}

	return true
}

func AllDecreasing[T cmp.Ordered](in []T) bool {
	for i := 1; i < len(in); i++ {
		if in[i] > in[i-1] {
			return false
		}
	}

	return true
}

// SplitToInts splits s with separator s and return a slice of int.
func SplitToInts(s, sep string) []int {
	parts := strings.Split(s, sep)
	result := make([]int, len(parts))
	for i, part := range parts {
		result[i] = MustAtoi(part)
	}

	return result
}

// ConcatRunes replaces consecutive occurrences of r in s with a single r, up to max occurrences.
// Absolutely not efficient.
func ConcatRunes(s string, r string, max int) string {
	for i := max; i >= 2; i-- {
		s = strings.ReplaceAll(s, strings.Repeat(r, i), r)
	}
	return s
}

// SplitInTwoInts splits s with separator s and return the 2 ints.
func SplitInTwoInts(s, sep string) (int, int) {
	parts := strings.Split(s, sep)
	if len(parts) != 2 {
		panic("expected exactly 2 parts")
	}

	return MustAtoi(parts[0]), MustAtoi(parts[1])
}

// CountToMap counts the number of occurrences for each element of in.
func CountToMap[T comparable](in []T) map[T]int {
	m := make(map[T]int)
	for _, n := range in {
		m[n] = m[n] + 1
	}

	return m
}

func Count[T comparable](slice []T, val T) int {
	count := 0
	for _, s := range slice {
		if s == val {
			count++
		}
	}
	return count
}

func Lines(s string) []string {
	lines := strings.Split(s, "\n")
	if lines[len(lines)-1] == "" {
		return lines[:len(lines)-1]
	}
	return lines
}

func LinesTwoParts(s string) ([]string, []string) {
	lines := strings.Split(s, "\n")
	ind := slices.Index(lines, "")

	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	return lines[:ind], lines[ind+1:]
}

// CountConsecutiveBackward count consecutive occurrences of val on slice s, starting at position start.
// Stops at the first non-equal value.
func CountConsecutiveBackward[T comparable](s []T, start int, val T) int {
	r := 0
	for i := start; i >= 0; i-- {
		if s[i] == val {
			r++
		} else {
			return r
		}
	}
	return r
}

// FindConsecutive tries to find a repetition of val of at least size before cutoff, returns the index of the first occurrence.
// returns -1 if not found.
func FindConsecutive[T comparable](s []T, size int, val T, cutoff int) int {
	free := 0
	for i, v := range s {
		if i == cutoff {
			return -1
		}
		if v == val {
			free++
			if free == size {
				return i - size + 1
			}
		} else {
			free = 0
		}
	}

	return -1
}

func HighestNumForward(s string) (val int, idx int) {
	val = -1
	for i := 0; i < len(s); i++ {
		n := MustAtoi(string(s[i]))

		if n > val {
			val = n
			idx = i
		}
	}
	return val, idx
}

func HighestNumBackward(s string) (val int, idx int) {
	val = -1
	for i := len(s) - 1; i >= 0; i-- {
		n := MustAtoi(string(s[i]))

		if n > val {
			val = n
			idx = i
		}
	}
	return val, idx
}
