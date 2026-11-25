class Solution:
    def smallestUniqueSubarray(self, nums: List[int]) -> int:
        n = len(nums)
        MOD1, MOD2 = 10**9 + 7, 10**9 + 9
        BASE1, BASE2 = 1000003, 1000033
        h1 = [0] * (n + 1)
        h2 = [0] * (n + 1)
        p1 = [1] * (n + 1)
        p2 = [1] * (n + 1)
        for i in range(n):
            h1[i + 1] = (h1[i] * BASE1 + nums[i]) % MOD1
            h2[i + 1] = (h2[i] * BASE2 + nums[i]) % MOD2
            p1[i + 1] = p1[i] * BASE1 % MOD1
            p2[i + 1] = p2[i] * BASE2 % MOD2
        def has_unique(length):
            cnt = {}
            for i in range(n - length + 1):
                a = (h1[i + length] - h1[i] * p1[length] % MOD1 + MOD1) % MOD1
                b = (h2[i + length] - h2[i] * p2[length] % MOD2 + MOD2) % MOD2
                key = (a, b)
                cnt[key] = cnt.get(key, 0) + 1
            return any(v == 1 for v in cnt.values())
        lo, hi, ans = 1, n, n
        while lo <= hi:
            mid = lo + (hi - lo) // 2
            if has_unique(mid):
                ans = mid
                hi = mid - 1
            else:
                lo = mid + 1
        return ans