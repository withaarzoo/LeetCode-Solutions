class Solution:
    def minimumChanges(self, s: str, k: int) -> int:
        n = len(s)
        INF = 10**9
        cost = [[INF] * n for _ in range(n)]
        for i in range(n):
            for j in range(i + 1, n):
                L = j - i + 1
                best = INF
                for d in range(1, L):
                    if L % d != 0:
                        continue
                    cur = 0
                    m = L // d
                    for r in range(d):
                        for t in range(m // 2):
                            if s[i + r + t * d] != s[i + r + (m - 1 - t) * d]:
                                cur += 1
                    best = min(best, cur)
                cost[i][j] = best
        dp = [[INF] * (k + 1) for _ in range(n + 1)]
        dp[0][0] = 0
        for i in range(1, n + 1):
            for p in range(1, k + 1):
                for prev in range(i - 1):
                    if prev <= i - 2 and dp[prev][p - 1] != INF and cost[prev][i - 1] != INF:
                        dp[i][p] = min(dp[i][p], dp[prev][p - 1] + cost[prev][i - 1])
        return dp[n][k]