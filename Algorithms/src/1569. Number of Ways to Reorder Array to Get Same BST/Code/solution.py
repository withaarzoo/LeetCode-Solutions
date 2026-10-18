class Solution:
    def numOfWays(self, nums: list[int]) -> int:
        MOD = 10**9 + 7
        n = len(nums)
        fact = [1] * (n + 1)
        for i in range(1, n + 1):
            fact[i] = fact[i - 1] * i % MOD
        inv = [1] * (n + 1)
        inv[n] = pow(fact[n], MOD - 2, MOD)
        for i in range(n - 1, -1, -1):
            inv[i] = inv[i + 1] * (i + 1) % MOD
        def comb(a, b):
            if b < 0 or b > a:
                return 0
            return fact[a] * inv[b] % MOD * inv[a - b] % MOD
        def ways(arr):
            m = len(arr)
            if m <= 2:
                return 1
            root = arr[0]
            left = [x for x in arr if x < root]
            right = [x for x in arr if x > root]
            return comb(m - 1, len(left)) * ways(left) % MOD * ways(right) % MOD
        return (ways(nums) - 1) % MOD