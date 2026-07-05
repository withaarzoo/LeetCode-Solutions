/**
 * @param {string[]} board
 * @return {number[]}
 */
var pathsWithMaxScore = function (board) {
  const MOD = 1000000007;
  const n = board.length;

  // These arrays store DP values for the row below.
  // -1 means that a cell cannot reach S.
  let nextScore = new Array(n + 1).fill(-1);
  let nextWays = new Array(n + 1).fill(0);

  // I process rows from bottom to top.
  for (let i = n - 1; i >= 0; i--) {
    // Fresh arrays hold the current row.
    const currScore = new Array(n + 1).fill(-1);
    const currWays = new Array(n + 1).fill(0);

    // I process right to left so the right cell is ready.
    for (let j = n - 1; j >= 0; j--) {
      const cell = board[i][j];

      // Obstacles are never reachable.
      if (cell === "X") {
        continue;
      }

      // S is the base case of the reversed DP.
      if (cell === "S") {
        currScore[j] = 0;
        currWays[j] = 1;
        continue;
      }

      // Find the best score from down, right, or diagonal.
      const best = Math.max(nextScore[j], currScore[j + 1], nextScore[j + 1]);

      // This cell cannot reach S if all next cells are unreachable.
      if (best === -1) {
        continue;
      }

      let ways = 0;

      // Add only paths that continue with the maximum score.
      if (nextScore[j] === best) {
        ways = (ways + nextWays[j]) % MOD;
      }
      if (currScore[j + 1] === best) {
        ways = (ways + currWays[j + 1]) % MOD;
      }
      if (nextScore[j + 1] === best) {
        ways = (ways + nextWays[j + 1]) % MOD;
      }

      // E adds 0; digit cells add their numeric value.
      const value = cell === "E" ? 0 : Number(cell);

      currScore[j] = best + value;
      currWays[j] = ways;
    }

    // The current row becomes the row below.
    nextScore = currScore;
    nextWays = currWays;
  }

  // No reachable value at E means no valid path exists.
  if (nextScore[0] === -1) {
    return [0, 0];
  }

  return [nextScore[0], nextWays[0]];
};
