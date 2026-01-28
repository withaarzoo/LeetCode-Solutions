package main
import (
	"sort"
	"math"
)

func min(a,b int64) int64 { if a<b { return a }; return b }

func minCost(grid [][]int, k int) int {
	m := len(grid)
	n := len(grid[0])
	const INF int64 = 4e18

	// dp base (no teleport)
	dp := make([][]int64, m)
	for i:=0;i<m;i++ {
		dp[i] = make([]int64, n)
		for j:=0;j<n;j++ { dp[i][j] = INF }
	}
	dp[0][0] = 0
	for i:=0;i<m;i++ {
		for j:=0;j<n;j++ {
			if i>0 {
				val := dp[i-1][j] + int64(grid[i][j])
				if val < dp[i][j] { dp[i][j] = val }
			}
			if j>0 {
				val := dp[i][j-1] + int64(grid[i][j])
				if val < dp[i][j] { dp[i][j] = val }
			}
		}
	}

	// prepare cells list sorted by value desc
	type Cell struct { v, i, j int }
	cells := make([]Cell, 0, m*n)
	for i:=0;i<m;i++ {
		for j:=0;j<n;j++ {
			cells = append(cells, Cell{grid[i][j], i, j})
		}
	}
	sort.Slice(cells, func(a,b int) bool { return cells[a].v > cells[b].v })

	for step:=0; step<k; step++ {
		start := make([][]int64, m)
		for i:=0;i<m;i++ {
			start[i] = make([]int64, n)
			for j:=0;j<n;j++ { start[i][j] = INF }
		}
		runningMin := INF
		idx := 0
		for idx < len(cells) {
			val := cells[idx].v
			j := idx
			minGroup := INF
			for j < len(cells) && cells[j].v == val {
				ii := cells[j].i; jj := cells[j].j
				if dp[ii][jj] < minGroup { minGroup = dp[ii][jj] }
				j++
			}
			if minGroup < runningMin { runningMin = minGroup }
			for p := idx; p < j; p++ {
				ii := cells[p].i; jj := cells[p].j
				if dp[ii][jj] < runningMin { start[ii][jj] = dp[ii][jj] } else { start[ii][jj] = runningMin }
			}
			idx = j
		}

		dp2 := make([][]int64, m)
		for i:=0;i<m;i++ {
			dp2[i] = make([]int64, n)
			for j:=0;j<n;j++ { dp2[i][j] = INF }
		}
		for i:=0;i<m;i++ {
			for j:=0;j<n;j++ {
				if start[i][j] < dp2[i][j] { dp2[i][j] = start[i][j] }
				if dp2[i][j] < INF {
					if i+1 < m {
						v := dp2[i][j] + int64(grid[i+1][j])
						if v < dp2[i+1][j] { dp2[i+1][j] = v }
					}
					if j+1 < n {
						v := dp2[i][j] + int64(grid[i][j+1])
						if v < dp2[i][j+1] { dp2[i][j+1] = v }
					}
				}
			}
		}
		dp = dp2
	}

	return int(dp[m-1][n-1])
}
