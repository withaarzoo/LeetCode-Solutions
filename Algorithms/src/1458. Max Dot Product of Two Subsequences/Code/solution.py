class Solution:
    def maxDotProduct(self, nums1: List[int], nums2: List[int]) -> int:
        n, m = len(nums1), len(nums2)
        
        dp = [[float('-inf')] * (m + 1) for _ in range(n + 1)]
        
        for i in range(n - 1, -1, -1):
            for j in range(m - 1, -1, -1):
                product = nums1[i] * nums2[j]
                
                take_both = product
                if dp[i + 1][j + 1] != float('-inf'):
                    take_both = max(take_both, product + dp[i + 1][j + 1])
                
                dp[i][j] = max(
                    take_both,
                    dp[i + 1][j],
                    dp[i][j + 1]
                )
        
        return dp[0][0]
