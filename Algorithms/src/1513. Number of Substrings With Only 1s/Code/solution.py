class Solution:
    def numSub(self, s: str) -> int:
        MOD = 10**9 + 7
        res = 0
        cnt = 0  # current consecutive '1's count

        for ch in s:
            if ch == '1':
                cnt += 1
            else:
                res = (res + (cnt * (cnt + 1) // 2) % MOD) % MOD
                cnt = 0

        # add last block if any
        res = (res + (cnt * (cnt + 1) // 2) % MOD) % MOD
        return res
