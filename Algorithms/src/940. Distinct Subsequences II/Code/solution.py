class Solution:
    def distinctSubseqII(self, s: str) -> int:
        MOD = 1000000007  # I use this modulo because the answer can be very large.
        
        dp = 1  # I start with the empty subsequence as the only subsequence.
        last = [0] * 26  # last[c] stores dp from before the previous occurrence of character c.
        
        for ch in s:  # I process every character exactly once.
            index = ord(ch) - ord('a')  # I convert the lowercase character into an index from 0 to 25.
            
            old_dp = dp  # I save the old count before changing dp.
            
            dp = (2 * dp - last[index] + MOD) % MOD  # I double the count and remove duplicate subsequences.
            
            last[index] = old_dp  # I store the old count for the next occurrence of this character.
        
        return (dp - 1 + MOD) % MOD  # I remove the empty subsequence from the final answer.