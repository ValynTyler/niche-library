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

	div_imp(books[0:10])
}

func div_imp(v []string) {
	length := len(v)
	middle := length / 2

	if length > 1 {
		left := v[0:middle]
		right := v[middle:length]

		div_imp(left)
		div_imp(right)
	} else {
		fmt.Println(v)
	}
}
