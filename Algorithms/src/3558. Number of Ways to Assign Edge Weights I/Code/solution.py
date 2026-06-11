class Solution:
    def assignEdgeWeights(self, edges: List[List[int]]) -> int:
        MOD = 10**9 + 7
        n = len(edges) + 1

        # Build adjacency list
        graph = [[] for _ in range(n + 1)]

        for u, v in edges:
            graph[u].append(v)
            graph[v].append(u)

        max_depth = 0

        # Iterative DFS -> (node, depth)
        stack = [(1, 0)]
        visited = [False] * (n + 1)
        visited[1] = True

        while stack:
            node, depth = stack.pop()

            max_depth = max(max_depth, depth)

            for nxt in graph[node]:
                if not visited[nxt]:
                    visited[nxt] = True
                    stack.append((nxt, depth + 1))

        # Number of valid assignments = 2^(max_depth - 1)
        return pow(2, max_depth - 1, MOD)