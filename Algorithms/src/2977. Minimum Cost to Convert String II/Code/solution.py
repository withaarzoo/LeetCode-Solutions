class Solution:
    def minimumCost(self, source, target, original, changed, cost):
        INF = 10**30
        id = {}
        lens = set()
        sz = 0

        dist = [[INF]*201 for _ in range(201)]

        for s, t, c in zip(original, changed, cost):
            if s not in id:
                id[s] = sz
                lens.add(len(s))
                sz += 1
            if t not in id:
                id[t] = sz
                sz += 1
            dist[id[s]][id[t]] = min(dist[id[s]][id[t]], c)

        for i in range(sz):
            dist[i][i] = 0

        for k in range(sz):
            for i in range(sz):
                if dist[i][k] < INF:
                    for j in range(sz):
                        if dist[k][j] < INF:
                            dist[i][j] = min(dist[i][j], dist[i][k] + dist[k][j])

        n = len(source)
        dp = [INF] * (n + 1)
        dp[0] = 0

        for i in range(n):
            if dp[i] == INF:
                continue

            if source[i] == target[i]:
                dp[i + 1] = min(dp[i + 1], dp[i])

            for L in lens:
                if i + L > n:
                    continue
                s = source[i:i+L]
                t = target[i:i+L]
                if s in id and t in id:
                    dp[i + L] = min(dp[i + L], dp[i] + dist[id[s]][id[t]])

        return -1 if dp[n] == INF else dp[n]
