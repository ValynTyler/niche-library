# Niche Library
A simple go script that generates all possible permutations of a Technical Computer Science student's bookshelf, given a set of organization rules.

The bookshelf is meant to hold ten books, and they are as follows:
- “Calculus A” by Adam Zaber,
- “Calculus B” by Bruce Yonder,
- “Linear Algebra” by Clive Saber,
- “Probability Theory” by David Faber,
- “Discrete Mathematics” by Eric Tander,
- “Statistics” by Fred Pater,
- “Algorithms and Complexity”  by Greg Norer,
- “Languages and Machines” by Hugh Qatar,
- “Algorithmic Discrete Mathematics” by Ian Mirror,
- “Graph Isomorphisms” by Jack Loner,

Given this student's particular aesthetic prefferences, the set of possible sollutions must definitively **exclude** cases where:
1. “Calculus A” is **not** on the left of “Calculus B”
2. “Algorithmic Discrete Mathematics” is **not** on the right of “Discrete Mathematics”
3. Eric Tander’s “Discrete Mathematics” **isn't** immediately followed by Fred Pater’s “Statistics”

Given these specification, my model generates 90.720 possible valid cases, which is hopefully the actual right solution to this problem.
