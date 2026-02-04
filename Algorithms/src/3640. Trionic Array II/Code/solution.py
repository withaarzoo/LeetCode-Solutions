class Solution:
    def maxSumTrionic(self, nums):
        n = len(nums)

        left = [0]*n
        right = [0]*n

        for i in range(n):
            left[i] = nums[i]
            if i > 0 and nums[i-1] < nums[i] and left[i-1] > 0:
                left[i] += left[i-1]

        for i in range(n-1, -1, -1):
            right[i] = nums[i]
            if i+1 < n and nums[i] < nums[i+1] and right[i+1] > 0:
                right[i] += right[i+1]

        parts = []
        l, s = 0, nums[0]
        for i in range(1, n):
            if nums[i-1] <= nums[i]:
                parts.append((l, i-1, s))
                l, s = i, 0
            s += nums[i]
        parts.append((l, n-1, s))

        ans = -10**30
        for p, q, s in parts:
            if p > 0 and q < n-1 and nums[p-1] < nums[p] and nums[q] < nums[q+1] and p < q:
                ans = max(ans, left[p-1] + s + right[q+1])
        return ans
