package main

import (
	"fmt"
	"slices"
)

func main() {
	books := [10]string{
		/* 0 */ "Adam Zaber, “Calculus A”",
		/* 1 */ "Bruce Yonder, “Calculus B”",
		/* 2 */ "Clive Saber, “Linear Algebra”",
		/* 3 */ "David Faber, “Probability Theory”",
		/* 4 */ "Eric Tander, “Discrete Mathematics”",
		/* 5 */ "Fred Pater, “Statistics”",
		/* 6 */ "Greg Norer, “Algorithms and Complexity” ",
		/* 7 */ "Hugh Qatar, “Languages and Machines”",
		/* 8 */ "Ian Mirror, “Algorithmic Discrete Mathematics”",
		/* 9 */ "Jack Loner, “Graph Isomorphisms”",
	}

	/*
		Rules:
		- [x] 0 is left of 1
		- [x] 4 is left of 8
		- [x] 4 is immediately followed by 5
	*/

	solution := [10]int{}

	count := backtrack(books[:], solution[:], 0)
	fmt.Println(count)
}

func backtrack(v []string, x []int, k int) int {
	count := 0
	for i := range v {
		x[k] = i
		if check(x, k) {
			if solves(v, k) {
				fmt.Println(x)
				count += 1
			} else {
				count += backtrack(v, x, k+1)
			}
		}
	}
	return count
}

func check(x []int, k int) bool {
	// distinct elements
	if slices.Contains(x[:k], x[k]) {
		return false
	}
	// rule I
	if x[k] == 1 && !slices.Contains(x[:k], 0) {
		return false
	}
	// rule II
	if x[k] == 8 && !slices.Contains(x[:k], 4) {
		return false
	}
	// rule III
	if (x[k] == 5 && k == 0) || (x[k] == 5 && x[k-1] != 4) {
		return false
	}

	return true
}

func solves(v []string, k int) bool {
	return k == len(v)-1
}
