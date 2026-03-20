import "sort"

func minAbsDiff(grid [][]int, k int) [][]int {
    m := len(grid)
    n := len(grid[0])
    ans := make([][]int, m-k+1)
    for i := range ans {
        ans[i] = make([]int, n-k+1)
    }

    for i := 0; i+k <= m; i++ {
        for j := 0; j+k <= n; j++ {
            vals := make([]int, 0, k*k)

            // Collect all values from the current k x k submatrix
            for r := i; r < i+k; r++ {
                for c := j; c < j+k; c++ {
                    vals = append(vals, grid[r][c])
                }
            }

            sort.Ints(vals)

            best := int(1 << 60)

            // Check only consecutive different values
            for x := 1; x < len(vals); x++ {
                if vals[x] != vals[x-1] {
                    diff := vals[x] - vals[x-1]
                    if diff < best {
                        best = diff
                    }
                }
            }

            if best == int(1<<60) {
                ans[i][j] = 0
            } else {
                ans[i][j] = best
            }
        }
    }

    return ans
}