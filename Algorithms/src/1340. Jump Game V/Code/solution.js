/**
 * @param {number[]} arr
 * @param {number} d
 * @return {number}
 */
var maxJumps = function (arr, d) {
  const n = arr.length;

  // dp[i] stores maximum jumps starting from i
  const dp = new Array(n).fill(-1);

  // DFS function
  const dfs = (i) => {
    // Return stored answer
    if (dp[i] !== -1) return dp[i];

    // Current index counts as 1
    let ans = 1;

    // Move right
    for (let j = i + 1; j <= Math.min(n - 1, i + d); j++) {
      // Stop if blocked
      if (arr[j] >= arr[i]) break;

      // Update answer
      ans = Math.max(ans, 1 + dfs(j));
    }

    // Move left
    for (let j = i - 1; j >= Math.max(0, i - d); j--) {
      // Stop if blocked
      if (arr[j] >= arr[i]) break;

      // Update answer
      ans = Math.max(ans, 1 + dfs(j));
    }

    // Store answer
    dp[i] = ans;

    return ans;
  };

  let answer = 1;

  // Start from every index
  for (let i = 0; i < n; i++) {
    answer = Math.max(answer, dfs(i));
  }

  return answer;
};
