func pathsWithMaxScore(board []string) []int {
    const MOD int = 1000000007
    n := len(board)

    // These arrays store DP values for the row below.
    nextScore := make([]int, n+1)
    nextWays := make([]int, n+1)

    // -1 means that a cell cannot reach S.
    for j := 0; j <= n; j++ {
        nextScore[j] = -1
    }

    // I process rows from bottom to top.
    for i := n - 1; i >= 0; i-- {
        // Fresh arrays store the current row.
        currScore := make([]int, n+1)
        currWays := make([]int, n+1)

        // Every score starts as unreachable.
        for j := 0; j <= n; j++ {
            currScore[j] = -1
        }

        // I process right to left so the right cell is already solved.
        for j := n - 1; j >= 0; j-- {
            cell := board[i][j]

            // Obstacles cannot be used.
            if cell == 'X' {
                continue
            }

            // S is the base case of the reversed DP.
            if cell == 'S' {
                currScore[j] = 0
                currWays[j] = 1
                continue
            }

            // Find the best score among down, right, and diagonal.
            best := nextScore[j]
            if currScore[j+1] > best {
                best = currScore[j+1]
            }
            if nextScore[j+1] > best {
                best = nextScore[j+1]
            }

            // No reachable next cell means this cell is unreachable.
            if best == -1 {
                continue
            }

            ways := 0

            // Count every next cell that gives the best score.
            if nextScore[j] == best {
                ways = (ways + nextWays[j]) % MOD
            }
            if currScore[j+1] == best {
                ways = (ways + currWays[j+1]) % MOD
            }
            if nextScore[j+1] == best {
                ways = (ways + nextWays[j+1]) % MOD
            }

            // E adds no points; digit cells add their digit value.
            value := 0
            if cell != 'E' {
                value = int(cell - '0')
            }

            currScore[j] = best + value
            currWays[j] = ways
        }

        // The current row becomes the row below.
        nextScore = currScore
        nextWays = currWays
    }

    // An unreachable E means no valid path exists.
    if nextScore[0] == -1 {
        return []int{0, 0}
    }

    return []int{nextScore[0], nextWays[0]}
}