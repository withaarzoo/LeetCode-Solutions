class Solution:
    def maxTotalValue(self, nums: List[int], k: int) -> int:
        # Global minimum element
        mn = min(nums)

        # Global maximum element
        mx = max(nums)

        # Best subarray value
        best = mx - mn

        # Choose the same best subarray k times
        return best * k