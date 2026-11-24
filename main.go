package main

import "fmt"

// MinPathSum returns the cheapest top-left to bottom-right path cost moving
// only right or down, and one cheapest path's coordinates, recovered by
// walking the DP table backwards.
func MinPathSum(grid [][]int) (int, [][2]int) {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0, nil
	}
	rows, cols := len(grid), len(grid[0])
	dp := make([][]int, rows)
	for r := range dp {
		dp[r] = make([]int, cols)
		for c := range dp[r] {
			switch {
			case r == 0 && c == 0:
				dp[r][c] = grid[0][0]
			case r == 0:
				dp[r][c] = dp[r][c-1] + grid[r][c]
			case c == 0:
				dp[r][c] = dp[r-1][c] + grid[r][c]
			default:
				dp[r][c] = min(dp[r-1][c], dp[r][c-1]) + grid[r][c]
			}
		}
	}
	// Backtrack from the goal: at each cell, step to the cheaper predecessor.
	path := [][2]int{}
	r, c := rows-1, cols-1
	for {
		path = append(path, [2]int{r, c})
		if r == 0 && c == 0 {
			break
		}
		if r == 0 {
			c--
		} else if c == 0 {
			r--
		} else if dp[r-1][c] <= dp[r][c-1] {
			r--
		} else {
			c--
		}
	}
	for l, x := 0, len(path)-1; l < x; l, x = l+1, x-1 {
		path[l], path[x] = path[x], path[l]
	}
	return dp[rows-1][cols-1], path
}

func main() {
	cost, path := MinPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}})
	fmt.Println(cost, path)
}
