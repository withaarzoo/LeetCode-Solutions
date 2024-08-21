/**
 * @param {string} s - The input string that the strange printer needs to print.
 * @return {number} - The minimum number of turns the printer needs to print the entire string.
 */
var strangePrinter = function (s) {
  const n = s.length; // Get the length of the input string

  // Initialize a 2D DP array with dimensions n x n and fill it with zeros.
  // dp[i][j] will store the minimum number of turns required to print the substring s[i:j+1].
  const dp = Array.from({ length: n }, () => Array(n).fill(0));

  // Iterate over the string in reverse order to fill the DP table
  for (let i = n - 1; i >= 0; i--) {
    dp[i][i] = 1; // Base case: A single character needs 1 turn to print

    // Consider all substrings starting at index i and ending at index j
    for (let j = i + 1; j < n; j++) {
      // Assume we print the character s[j] separately, hence adding 1 turn
      dp[i][j] = dp[i][j - 1] + 1;

      // Check for possible combinations to minimize the turns
      // Iterate over possible partitions within the substring s[i:j+1]
      for (let k = i; k < j; k++) {
        // If s[k] equals s[j], we can merge the turns
        if (s[k] === s[j]) {
          // Update dp[i][j] with the minimum turns between the current value and the value obtained by merging
          dp[i][j] = Math.min(dp[i][j], dp[i][k] + dp[k + 1][j - 1]);
        }
      }
    }
  }

  // The result is stored in dp[0][n-1], which represents the minimum turns required to print the entire string s.
  return dp[0][n - 1];
};
