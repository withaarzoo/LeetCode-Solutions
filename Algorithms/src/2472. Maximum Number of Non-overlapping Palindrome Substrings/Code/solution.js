/**
 * @param {string} s
 * @param {number} k
 * @return {number}
 */
var maxPalindromes = function(s, k) {
    const n = s.length;
    // isPal[i][j] == true  iff  s[i..j] is a palindrome
    const isPal = Array.from({length: n}, () => Array(n).fill(false));
    
    // every single character is a palindrome
    for (let i = 0; i < n; ++i)
        isPal[i][i] = true;
    
    // every pair of identical characters is a palindrome
    for (let i = 0; i + 1 < n; ++i)
        if (s[i] === s[i + 1])
            isPal[i][i + 1] = true;
    
    // longer lengths: ends equal and inner part already known to be palindrome
    for (let len = 3; len <= n; ++len)
        for (let i = 0; i + len - 1 < n; ++i) {
            const j = i + len - 1;
            if (s[i] === s[j] && isPal[i + 1][j - 1])
                isPal[i][j] = true;
        }
    
    // dp[i] = maximum number of valid pieces inside s[0..i-1]
    const dp = new Array(n + 1).fill(0);
    for (let i = 1; i <= n; ++i) {
        dp[i] = dp[i - 1];                 // skip the last character
        // try every possible start of a piece that ends at i-1
        for (let j = 0; j <= i - k; ++j)
            if (isPal[j][i - 1])
                dp[i] = Math.max(dp[i], dp[j] + 1);
    }
    return dp[n];
};