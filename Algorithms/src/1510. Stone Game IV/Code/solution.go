func winnerSquareGame(n int) bool {
	// dp[i] tells me whether the current player can win with i stones.
	dp := make([]bool, n+1)

	// With 0 stones, the current player has no move and loses.
	dp[0] = false

	// Calculate the result for every number of stones from 1 to n.
	for i := 1; i <= n; i++ {
		// Try every perfect square that can be removed from i.
		for j := 1; j*j <= i; j++ {
			// If the remaining state is losing for the opponent,
			// I can remove this square and force a win.
			if !dp[i-j*j] {
				dp[i] = true

				// One winning move is enough, so stop checking squares.
				break
			}
		}
	}

	// Return whether Alice can force a win with n stones.
	return dp[n]
}