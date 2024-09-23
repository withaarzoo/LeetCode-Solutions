var minExtraChar = function (s, dictionary) {
  // Create a set from the dictionary array for O(1) time complexity when checking if a word exists
  let dict = new Set(dictionary);

  // Get the length of the input string `s`
  let n = s.length;

  // Initialize a DP (Dynamic Programming) array of size `n + 1` (for zero-indexing),
  // and fill it with the value `n` (maximum possible extra characters).
  // `dp[i]` will store the minimum extra characters required for the substring `s[0:i]`
  let dp = new Array(n + 1).fill(n);

  // Base case: If the string is empty (s = ""), no extra characters are required
  dp[0] = 0;

  // Iterate through the string, considering substrings from the start to each position `i`
  for (let i = 1; i <= n; i++) {
    // Explore all substrings that end at position `i` by varying the start position `j`
    for (let j = 0; j < i; j++) {
      // Extract the substring `s[j:i]`
      let sub = s.substring(j, i);

      // Check if the current substring `sub` exists in the dictionary
      if (dict.has(sub)) {
        // If it does, update `dp[i]` to the minimum of its current value
        // or the value at `dp[j]` (the number of extra characters for the substring `s[0:j]`)
        dp[i] = Math.min(dp[i], dp[j]);
      }
    }
    // After trying all substrings ending at `i`, consider the current character `s[i-1]`
    // as an extra character, and update `dp[i]` accordingly.
    dp[i] = Math.min(dp[i], dp[i - 1] + 1);
  }

  // The result is stored in `dp[n]`, which holds the minimum extra characters for the entire string `s`
  return dp[n];
};
