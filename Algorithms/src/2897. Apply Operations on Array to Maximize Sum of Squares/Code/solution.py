class Solution:
    def maxSum(self, nums: List[int], k: int) -> int:
        MOD = 10**9 + 7
        n = len(nums)
        cnt = [0] * 30
        for x in nums:
            for b in range(30):
                if x & (1 << b):
                    cnt[b] += 1
        a = [0] * n
        for b in range(29, -1, -1):
            c = cnt[b]
            for i in range(c):
                a[i] |= (1 << b)
        ans = 0
        for i in range(k):
            ans = (ans + a[i] * a[i]) % MOD
        return ans