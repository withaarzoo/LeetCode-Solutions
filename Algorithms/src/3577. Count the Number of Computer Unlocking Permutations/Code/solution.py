class Solution:
    def countPermutations(self, complexity: List[int]) -> int:
        MOD = 10**9 + 7
        n = len(complexity)
        
        # 1. Find global minimum and its count
        min_val = min(complexity)
        cnt_min = complexity.count(min_val)
        
        # 2. Check if index 0 has unique minimum
        if complexity[0] != min_val or cnt_min != 1:
            return 0
        
        # 3. Compute (n - 1)! % MOD
        ans = 1
        for i in range(2, n):
            ans = (ans * i) % MOD
        return ans
