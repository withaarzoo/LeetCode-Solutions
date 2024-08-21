class Solution:
    def strangePrinter(self, s: str) -> int:
        # Step 1: Calculate the length of the string `s`
        n = len(s)
        
        # Step 2: Initialize a 2D DP table with dimensions n x n
        # dp[i][j] represents the minimum number of turns needed to print the substring s[i:j+1]
        dp = [[0] * n for _ in range(n)]
        
        # Step 3: Start filling the DP table from the bottom row to the top row
        for i in range(n - 1, -1, -1):
            # Step 4: Base case: A single character (i == j) needs exactly 1 turn to print
            dp[i][i] = 1
            
            # Step 5: Iterate over all possible substrings starting from i to j
            for j in range(i + 1, n):
                # Step 6: By default, assume we print s[j] separately, so we add 1 turn to the result
                dp[i][j] = dp[i][j - 1] + 1
                
                # Step 7: Check for all possible partitions of the substring s[i:j+1]
                for k in range(i, j):
                    # Step 8: If the characters at positions k and j are the same,
                    # it means we can potentially minimize the number of turns
                    # by merging the print jobs
                    if s[k] == s[j]:
                        dp[i][j] = min(dp[i][j], dp[i][k] + dp[k + 1][j - 1])
        
        # Step 9: The final result is stored in dp[0][n - 1], 
        # which gives the minimum number of turns needed to print the entire string s
        return dp[0][n - 1]
