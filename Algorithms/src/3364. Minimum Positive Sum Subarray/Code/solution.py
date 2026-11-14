class Solution:
    def minimumSumSubarray(self, nums: List[int], l: int, r: int) -> int:
        n = len(nums)
        ans = float('inf')
        for i in range(n):
            s = 0
            for j in range(i, n):
                s += nums[j]
                length = j - i + 1
                if l <= length <= r and s > 0:
                    ans = min(ans, s)
        return -1 if ans == float('inf') else ans