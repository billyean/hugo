package array

// LeetCode's exercise.
// https://leetcode.com/problems/count-negative-numbers-in-a-sorted-matrix/submissions/
func CountNegatives(grid [][]int) int {
	rz := len(grid[0])
	r := len(grid) - 1
	c := 0
	count := 0
	for r >= 0 && c < rz {
		if grid[r][c] < 0 {
			count += rz - c
			r--
		} else {
			c++
		}
	}

	return count
}
