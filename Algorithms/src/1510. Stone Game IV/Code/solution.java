class Solution {
    public boolean winnerSquareGame(int n) {
        // dp[i] tells me whether the current player can win with i stones.
        boolean[] dp = new boolean[n + 1];

        // With 0 stones, the current player has no valid move and loses.
        dp[0] = false;

        // Calculate the result for every number of stones from 1 to n.
        for (int i = 1; i <= n; i++) {
            // Try every perfect square that can be removed from i.
            for (int j = 1; j * j <= i; j++) {
                // If the remaining state is losing for the opponent,
                // this move lets me force a win.
                if (!dp[i - j * j]) {
                    dp[i] = true;

                    // I only need one winning move, so I can stop here.
                    break;
                }
            }
        }

        // Return whether Alice can win with n stones.
        return dp[n];
    }
}