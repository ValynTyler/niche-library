package main

import "fmt"

func main() {
	books := [10]string{
		"Adam Zaber, “Calculus A”",
		"Bruce Yonder, “Calculus B”",
		"Clive Saber, “Linear Algebra”",
		"David Faber, “Probability Theory”",
		"Eric Tander, “Discrete Mathematics”",
		"Fred Pater, “Statistics”",
		"Greg Norer, “Algorithms and Complexity” ",
		"Hugh Qatar, “Languages and Machines”",
		"Ian Mirror, “Algorithmic Discrete Mathematics”",
		"Jack Loner, “Graph Isomorphisms”",
	}

	solution := [10]int{}

	count := backtrack(books[:], solution[:], 0)
	fmt.Println(count)
}

func backtrack(v []string, x []int, k int) int {
	count := 0
	for i := range v {
		x[k] = i
		if check(v, x, k) {
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

func check(v []string, x []int, k int) bool {
	for i := range v[:k] {
		if x[k] == x[i] {
			return false
		}
	}
	return true
}

func solves(v []string, k int) bool {
	return k == len(v)-1
}
