class Solution:
    def xorAfterQueries(self, nums: List[int], queries: List[List[int]]) -> int:
        MOD = 10**9 + 7

        # Process each query
        for l, r, k, v in queries:
            # Visit indices: l, l+k, l+2k, ... <= r
            for i in range(l, r + 1, k):
                nums[i] = (nums[i] * v) % MOD

        # Compute XOR of all final values
        ans = 0
        for num in nums:
            ans ^= num

        return ans