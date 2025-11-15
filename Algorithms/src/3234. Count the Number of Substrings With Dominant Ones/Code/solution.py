class Solution:
    def numberOfSubstrings(self, s: str) -> int:
        n = len(s)
        ans = 0
        # all-ones substrings
        run = 0
        for ch in s:
            if ch == '1':
                run += 1
            else:
                ans += run * (run + 1) // 2
                run = 0
        ans += run * (run + 1) // 2

        zeroPos = [i for i, ch in enumerate(s) if ch == '0']
        m = len(zeroPos)
        if m == 0:
            return ans

        import math
        K = int(math.isqrt(n))
        for k in range(1, K + 1):
            if k > m: break
            for i in range(0, m - k + 1):
                leftPrev = -1 if i == 0 else zeroPos[i - 1]
                rightNext = n if (i + k - 1 == m - 1) else zeroPos[i + k]
                leftOnes = zeroPos[i] - leftPrev - 1
                rightOnes = rightNext - zeroPos[i + k - 1] - 1
                baseLen = zeroPos[i + k - 1] - zeroPos[i] + 1
                needLen = k * k + k
                t = needLen - baseLen
                totalPairs = (leftOnes + 1) * (rightOnes + 1)
                if t <= 0:
                    ans += totalPairs
                    continue

                pairs_lt = 0
                s0 = t - 1
                if s0 >= 0:
                    L = leftOnes
                    R = rightOnes
                    x_max = min(L, s0)
                    if x_max >= 0:
                        x0 = max(0, s0 - R)
                        if x0 > x_max:
                            pairs_lt = (x_max + 1) * (R + 1)
                        else:
                            part1 = x0 * (R + 1)
                            n2 = x_max - x0 + 1
                            sum_x = (x0 + x_max) * n2 // 2
                            part2 = n2 * (s0 + 1) - sum_x
                            pairs_lt = part1 + part2
                    else:
                        pairs_lt = 0
                else:
                    pairs_lt = 0

                valid = totalPairs - pairs_lt
                if valid > 0:
                    ans += valid
        return ans
