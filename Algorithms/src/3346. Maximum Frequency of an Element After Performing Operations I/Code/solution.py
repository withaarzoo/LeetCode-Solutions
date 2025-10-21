class Solution:
    def maxFrequency(self, nums: List[int], k: int, numOperations: int) -> int:
        if not nums:
            return 0
        mx = max(nums)
        size = mx + k + 2
        count = [0] * size

        # frequency array
        for v in nums:
            count[v] += 1

        # prefix sums
        for i in range(1, size):
            count[i] += count[i-1]

        ans = 0
        for t in range(size):
            L = max(0, t - k)
            R = min(size - 1, t + k)
            total = count[R] - (count[L-1] if L > 0 else 0)   # elements that can become t
            freq_t = count[t] - (count[t-1] if t > 0 else 0) # how many already t
            canConvert = total - freq_t
            take = min(numOperations, canConvert)
            ans = max(ans, freq_t + take)
        return ans
