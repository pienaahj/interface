package main

import (
	"sort"
	"strings"
)

type list []*product

func (l list) String() string {
	if len(l) == 0 {

		return "Sorry.  We're waiting for delivery.\n "
	}

	var str strings.Builder
	for _, p := range l {
		// fmt.Printf("* %s\n", p)
		str.WriteString("* ")
		str.WriteString(p.String())
		str.WriteRune('\n')
	}
	return str.String()
}

func (l list) discount(ratio float64) {
	for _, p := range l {
		p.discount(ratio)
	}
}

// Len is the number of elements in the collection.
func (l list) Len() int {
	return len(l)
}

// Less reports whether the element with index i
// must sort before the element with index j.
//
// If both Less(i, j) and Less(j, i) are false,
// then the elements at index i and j are considered equal.
// Sort may place equal elements in any order in the final result,
// while Stable preserves the original input order of equal elements.
//
// Less must describe a transitive ordering:
//   - if both Less(i, j) and Less(j, k) are true, then Less(i, k) must be true as well.
//   - if both Less(i, j) and Less(j, k) are false, then Less(i, k) must be false as well.
//
// Note that floating-point comparison (the < operator on float32 or float64 values)
// is not a transitive ordering when not-a-number (NaN) values are involved.
// See Float64Slice.Less for a correct implementation for floating-point values.
func (l list) Less(i, j int) bool {
	return l[i].Title < l[j].Title
}

// Swap swaps the elements with indexes i and j.
func (l list) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

type byRelease struct {
	list
}

func (br byRelease) Less(i, j int) bool {
	return br.list[i].Released.Before(br.list[j].Released.Time)
}
func byReleaseDate(l list) sort.Interface {

	return &byRelease{l}
}
