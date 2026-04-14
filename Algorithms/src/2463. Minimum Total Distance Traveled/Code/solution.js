/**
 * @param {number[]} robot
 * @param {number[][]} factory
 * @return {number}
 */
var minimumTotalDistance = function (robot, factory) {
  robot.sort((a, b) => a - b);
  factory.sort((a, b) => a[0] - b[0]);

  const n = robot.length;
  const m = factory.length;
  const INF = Number.MAX_SAFE_INTEGER;

  const dp = Array.from({ length: n + 1 }, () => Array(m + 1).fill(-1));

  function solve(i, j) {
    // All robots repaired
    if (i === n) return 0;

    // No factories left
    if (j === m) return INF;

    if (dp[i][j] !== -1) return dp[i][j];

    // Skip current factory
    let ans = solve(i, j + 1);

    let distance = 0;
    const [pos, limit] = factory[j];

    // Use current factory for next k robots
    for (let k = 0; k < limit && i + k < n; k++) {
      distance += Math.abs(robot[i + k] - pos);

      const next = solve(i + k + 1, j + 1);

      if (next !== INF) {
        ans = Math.min(ans, distance + next);
      }
    }

    return (dp[i][j] = ans);
  }

  return solve(0, 0);
};
