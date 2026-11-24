package main

import "testing"

func TestMinPathSum(t *testing.T) {
	cost, path := MinPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}})
	if cost != 7 {
		t.Errorf("cost = %d, want 7", cost)
	}
	// The returned path must be valid and cost what was claimed.
	if path[0] != [2]int{0, 0} || path[len(path)-1] != [2]int{2, 2} {
		t.Errorf("path endpoints wrong: %v", path)
	}
	sum := 0
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	for i, p := range path {
		sum += grid[p[0]][p[1]]
		if i > 0 {
			dr := p[0] - path[i-1][0]
			dc := p[1] - path[i-1][1]
			if !(dr == 1 && dc == 0 || dr == 0 && dc == 1) {
				t.Errorf("illegal move %v -> %v", path[i-1], p)
			}
		}
	}
	if sum != cost {
		t.Errorf("path sums to %d, claimed %d", sum, cost)
	}

	if cost, _ := MinPathSum([][]int{{5}}); cost != 5 {
		t.Errorf("single cell = %d", cost)
	}
	if cost, _ := MinPathSum([][]int{{1, 2, 3}}); cost != 6 {
		t.Errorf("single row = %d", cost)
	}
	if cost, path := MinPathSum(nil); cost != 0 || path != nil {
		t.Error("empty grid")
	}
}
