class Solution:
    def remainingMethods(self, n: int, k: int, invocations: List[List[int]]) -> List[int]:

        # Build the graph
        graph = [[] for _ in range(n)]

        for u, v in invocations:
            graph[u].append(v)

        # Marks suspicious methods
        vis = [False] * n

        # DFS to visit every reachable method
        def dfs(u):
            vis[u] = True

            for v in graph[u]:
                if not vis[v]:
                    dfs(v)

        dfs(k)

        # If any safe method calls a suspicious method,
        # removal is impossible
        for u, v in invocations:
            if not vis[u] and vis[v]:
                return list(range(n))

        # Return only non-suspicious methods
        ans = []

        for i in range(n):
            if not vis[i]:
                ans.append(i)

        return ans