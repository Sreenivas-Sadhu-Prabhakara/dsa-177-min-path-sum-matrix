# Minimum Path Sum in a Matrix

> Striver SDE Sheet 177 · Day 26 DP II — each cell's best cost is its value plus the cheaper of its two predecessors; the sheet also asks for the path, recovered by walking the table backwards. Go, stdlib only, tested.

## Problem

Cheapest path from the top-left to the bottom-right of a cost grid, moving only right or down — and (per the sheet's phrasing) also *produce* a cheapest path, not just its cost.

## Approach

The DP is the two-predecessor minimum: `dp[r][c] = grid[r][c] + min(above, left)`, edges taking their only predecessor. The path-recovery half is the teachable part: don't store paths during the forward pass (quadratic memory of lists) — keep the finished table and **walk backwards** from the goal, at each cell stepping into whichever predecessor holds the smaller dp value. The backward walk re-derives the decisions for free. The test verifies the path as a property: legal moves only, endpoints right, and its cells summing to exactly the claimed cost.

**Complexity:** O(R·C) time; O(R·C) space (kept for the backtrack; cost-only would roll rows).

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
