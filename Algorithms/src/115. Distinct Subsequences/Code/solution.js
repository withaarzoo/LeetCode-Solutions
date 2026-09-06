/**
 * @param {string} s
 * @param {string} t
 * @return {number}
 */
var numDistinct = function (s, t) {
  const m = t.length;

  // dp[j] stores the number of ways to form the first j characters of t.
  // dp[0] is 1 because choosing nothing forms the empty string.
  const dp = new Array(m + 1).fill(0);
  dp[0] = 1;

  // Process each character of s.
  for (let i = 0; i < s.length; i++) {
    const c = s[i];

    // Traverse backwards so dp[j - 1] is not updated by the current character.
    // This makes sure one character of s is used only once in a transition.
    for (let j = m; j >= 1; j--) {
      // If the characters match, add the ways of forming the previous prefix.
      if (c === t[j - 1]) {
        dp[j] += dp[j - 1];
      }
    }
  }

  // dp[m] contains the number of distinct subsequences equal to t.
  return dp[m];
};
