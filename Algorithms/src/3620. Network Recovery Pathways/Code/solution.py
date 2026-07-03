class Solution:
    def findMaxPathScore(self, edges: List[List[int]], online: List[bool], k: int) -> int:
        n = len(online)

        # Build graph
        graph = [[] for _ in range(n)]
        indegree = [0] * n

        for u, v, w in edges:
            graph[u].append((v, w))
            indegree[v] += 1

        # Topological order
        from collections import deque

        q = deque()

        for i in range(n):
            if indegree[i] == 0:
                q.append(i)

        topo = []

        while q:
            u = q.popleft()
            topo.append(u)

            for v, _ in graph[u]:
                indegree[v] -= 1
                if indegree[v] == 0:
                    q.append(v)

        def check(limit):
            INF = 10 ** 30

            # Minimum cost to each node
            dp = [INF] * n
            dp[0] = 0

            for u in topo:

                # Skip unreachable nodes
                if dp[u] == INF:
                    continue

                # Skip offline intermediate nodes
                if u != 0 and u != n - 1 and not online[u]:
                    continue

                for v, w in graph[u]:

                    # Ignore small edges
                    if w < limit:
                        continue

                    # Cannot enter offline intermediate node
                    if v != n - 1 and not online[v]:
                        continue

                    if dp[u] + w < dp[v]:
                        dp[v] = dp[u] + w

            return dp[-1] <= k

        left = 0
        right = 10 ** 9
        ans = -1

        while left <= right:
            mid = (left + right) // 2

            if check(mid):
                ans = mid
                left = mid + 1
            else:
                right = mid - 1

        return ans