class Solution:
    def concatenatedBinary(self, n: int) -> int:
        MOD = 10**9 + 7
        ans = 0
        bit_length = 0
        
        for i in range(1, n + 1):
            
            # If i is power of 2
            if (i & (i - 1)) == 0:
                bit_length += 1
            
            # Shift and add
            ans = ((ans << bit_length) + i) % MOD
        
        return ans