class Solution:
    def maxPalindromes(self, s: str, k: int) -> int:
        n = len(s)
        # isPal[i][j] == True  iff  s[i..j] is a palindrome
        isPal = [[False] * n for _ in range(n)]
        
        # every single character is a palindrome
        for i in range(n):
            isPal[i][i] = True
        
        # every pair of identical characters is a palindrome
        for i in range(n - 1):
            if s[i] == s[i + 1]:
                isPal[i][i + 1] = True
        
        # longer lengths: ends equal and inner part already known to be palindrome
        for length in range(3, n + 1):
            for i in range(n - length + 1):
                j = i + length - 1
                if s[i] == s[j] and isPal[i + 1][j - 1]:
                    isPal[i][j] = True
        
        # dp[i] = maximum number of valid pieces inside s[0..i-1]
        dp = [0] * (n + 1)
        for i in range(1, n + 1):
            dp[i] = dp[i - 1]                 # skip the last character
            # try every possible start of a piece that ends at i-1
            for j in range(i - k + 1):
                if isPal[j][i - 1]:
                    dp[i] = max(dp[i], dp[j] + 1)
        return dp[n]