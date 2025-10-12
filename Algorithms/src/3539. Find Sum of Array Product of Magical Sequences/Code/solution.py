# Python3 implementation
from typing import List
MOD = 10**9 + 7

class Solution:
    def magicalSum(self, m: int, k: int, nums: List[int]) -> int:
        n = len(nums)
        # combinations
        C = [[0]*(m+1) for _ in range(m+1)]
        for i in range(m+1):
            C[i][0] = C[i][i] = 1
            for j in range(1,i):
                C[i][j] = (C[i-1][j-1] + C[i-1][j]) % MOD
        # powers
        powA = [[1]*(m+1) for _ in range(n)]
        for i in range(n):
            a = nums[i] % MOD
            for t in range(1, m+1):
                powA[i][t] = (powA[i][t-1] * a) % MOD
        M = m
        cur = [[[0]*(M+1) for _ in range(M+1)] for __ in range(M+1)]
        cur[M][0][0] = 1

        for i in range(n):
            nxt = [[[0]*(M+1) for _ in range(M+1)] for __ in range(M+1)]
            for r in range(M+1):
                for carry in range(M+1):
                    for ones in range(M+1):
                        val = cur[r][carry][ones]
                        if val == 0: continue
                        for t in range(r+1):
                            newr = r - t
                            s = carry + t
                            bit = s & 1
                            newones = ones + bit
                            if newones > M: continue
                            newcarry = s >> 1
                            mult = (C[r][t] * powA[i][t]) % MOD
                            add = (val * mult) % MOD
                            nxt[newr][newcarry][newones] = (nxt[newr][newcarry][newones] + add) % MOD
            cur = nxt

        ans = 0
        for carry in range(M+1):
            for ones in range(M+1):
                val = cur[0][carry][ones]
                if val == 0: continue
                extra = bin(carry).count("1")
                if ones + extra == k:
                    ans = (ans + val) % MOD
        return ans
