class Solution:
    def maxProfit(self, n, present, future, hierarchy, budget):
        tree = [[] for _ in range(n)]
        for u, v in hierarchy:
            tree[u - 1].append(v - 1)

        dp = [[[0] * (budget + 1) for _ in range(2)] for _ in range(n)]

        def merge(A, B):
            C = [-10**9] * (budget + 1)
            for i in range(budget + 1):
                if A[i] < 0: continue
                for j in range(budget - i + 1):
                    C[i + j] = max(C[i + j], A[i] + B[j])
            return C

        def dfs(u):
            for v in tree[u]:
                dfs(v)

            for parentBought in (0, 1):
                price = present[u] // 2 if parentBought else present[u]
                profit = future[u] - price

                skip = [0] * (budget + 1)
                for v in tree[u]:
                    skip = merge(skip, dp[v][0])

                take = [-10**9] * (budget + 1)
                if price <= budget:
                    base = [0] * (budget + 1)
                    for v in tree[u]:
                        base = merge(base, dp[v][1])
                    for b in range(price, budget + 1):
                        take[b] = base[b - price] + profit

                for b in range(budget + 1):
                    dp[u][parentBought][b] = max(skip[b], take[b])

        dfs(0)
        return max(dp[0][0])
