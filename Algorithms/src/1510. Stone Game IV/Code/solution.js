/**
 * @param {number} n
 * @return {boolean}
 */
var winnerSquareGame = function (n) {
  // dp[i] tells me whether the current player can win with i stones.
  const dp = new Array(n + 1).fill(false);

  // With 0 stones, there is no possible move, so the player loses.
  dp[0] = false;

  // Calculate the winning or losing state for every number of stones.
  for (let i = 1; i <= n; i++) {
    // Try removing every perfect square that is at most i.
    for (let j = 1; j * j <= i; j++) {
      // If the remaining state is losing for the opponent,
      // I can remove this square and force a win.
      if (!dp[i - j * j]) {
        dp[i] = true;

        // One winning move is enough, so stop checking.
        break;
      }
    }
  }

  // Return whether Alice can force a win from n stones.
  return dp[n];
};
