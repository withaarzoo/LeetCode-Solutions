class Solution:
    def minFlips(self, s: str) -> int:
        n = len(s)
        ss = s + s

        diff1 = 0
        diff2 = 0
        ans = float('inf')

        for i in range(len(ss)):

            expected1 = '0' if i % 2 == 0 else '1'
            expected2 = '1' if i % 2 == 0 else '0'

            if ss[i] != expected1:
                diff1 += 1
            if ss[i] != expected2:
                diff2 += 1

            if i >= n:
                prev = ss[i-n]

                prevExp1 = '0' if (i-n) % 2 == 0 else '1'
                prevExp2 = '1' if (i-n) % 2 == 0 else '0'

                if prev != prevExp1:
                    diff1 -= 1
                if prev != prevExp2:
                    diff2 -= 1

            if i >= n-1:
                ans = min(ans, diff1, diff2)

        return ans