package utils

import "fmt"

type rge struct {
	start int
	end   int
}

// DisjointRange manages a set of disjoint integer ranges.
// Not safe for concurrent use.
type DisjointRange struct {
	ranges []rge
}

// NewDisjointRange creates a new DisjointRange
func NewDisjointRange() *DisjointRange {
	return &DisjointRange{
		ranges: make([]rge, 0),
	}
}

func (dr *DisjointRange) AddRange(start, end int) {
	if start > end {
		panic(fmt.Sprintf("invalid range: start %d end %d", start, end))
	}
	startIndex := -1 // idx of existing range where start would fall into
	endIndex := -1   // idx of existing range where end would fall into
	// first we delete fully included ranges by making copies of ranges that are not fully included
	// it can probably be optimized to do in one pass but it would need some juggling with indices
	rgcopy := make([]rge, 0, len(dr.ranges))
	for _, r := range dr.ranges {
		if r.start < start || r.end > end {
			rgcopy = append(rgcopy, r)
		}
	}

	if len(rgcopy) != len(dr.ranges) {
		dr.ranges = rgcopy
	}

	for i, r := range dr.ranges {
		if start >= r.start && start <= r.end+1 { // +1 to allow touching ranges to merge
			startIndex = i
		}
		if end >= r.start-1 && end <= r.end { // -1 to allow touching ranges to merge
			endIndex = i
		}
	}

	idxToRemove := -1
	switch {
	case startIndex == -1 && endIndex == -1:
		// new separate range
		dr.ranges = append(dr.ranges, rge{start: start, end: end})
	case startIndex == endIndex:
		// both start and end fall into the same existing range, nothing to do
	case startIndex != -1 && endIndex != -1:
		// both start and end fall into different existing ranges, merge them
		dr.ranges[startIndex].end = dr.ranges[endIndex].end
		idxToRemove = endIndex
	case startIndex != -1:
		dr.ranges[startIndex].end = end
	default: //endIndex != -1
		dr.ranges[endIndex].start = start
	}
	if idxToRemove != -1 {
		dr.ranges = append(dr.ranges[:idxToRemove], dr.ranges[idxToRemove+1:]...)
	}
}

func (dr *DisjointRange) SumRanges() int {
	sum := 0
	for _, r := range dr.ranges {
		sum += r.end - r.start
	}
	return sum
}

// SumInclRanges returns the inclusive sum of all ranges
func (dr *DisjointRange) SumInclRanges() int {
	sum := 0
	for _, r := range dr.ranges {
		sum += r.end - r.start + 1
	}
	return sum
}

func (dr *DisjointRange) CountRanges() int {
	return len(dr.ranges)
}

// TODO: remove range
// TODO: isInOneRange
