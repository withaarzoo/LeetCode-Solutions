class Solution:
    def minOperations(self, nums: list[int], x: int) -> int:
        total = sum(nums)
        if total < x:
            return -1
        target = total - x
        if target == 0:
            return len(nums)
        n = len(nums)
        left = 0
        curr = 0
        max_len = -1
        for right in range(n):
            curr += nums[right]
            while curr > target and left <= right:
                curr -= nums[left]
                left += 1
            if curr == target:
                max_len = max(max_len, right - left + 1)
        return -1 if max_len == -1 else n - max_len