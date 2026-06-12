class Solution:
    def assignEdgeWeights(self, edges: List[List[int]], queries: List[List[int]]) -> List[int]:
        MOD = 1000000007

        n = len(edges) + 1

        LOG = 1
        while (1 << LOG) <= n:
            LOG += 1

        # Build adjacency list
        graph = [[] for _ in range(n + 1)]

        for u, v in edges:
            graph[u].append(v)
            graph[v].append(u)

        depth = [0] * (n + 1)
        up = [[1] * LOG for _ in range(n + 1)]

        # DFS preprocessing
        def dfs(node: int, parent: int) -> None:
            up[node][0] = parent

            for j in range(1, LOG):
                up[node][j] = up[up[node][j - 1]][j - 1]

            for nxt in graph[node]:
                if nxt == parent:
                    continue

                depth[nxt] = depth[node] + 1
                dfs(nxt, node)

        dfs(1, 1)

        # Binary lifting LCA
        def lca(a: int, b: int) -> int:
            if depth[a] < depth[b]:
                a, b = b, a

            diff = depth[a] - depth[b]

            for j in range(LOG - 1, -1, -1):
                if (diff >> j) & 1:
                    a = up[a][j]

            if a == b:
                return a

            for j in range(LOG - 1, -1, -1):
                if up[a][j] != up[b][j]:
                    a = up[a][j]
                    b = up[b][j]

            return up[a][0]

        # Precompute powers of two
        pow2 = [1] * (n + 1)

        for i in range(1, n + 1):
            pow2[i] = (pow2[i - 1] * 2) % MOD

        ans = []

        for u, v in queries:
            ancestor = lca(u, v)

            dist = depth[u] + depth[v] - 2 * depth[ancestor]

            if dist == 0:
                ans.append(0)
            else:
                ans.append(pow2[dist - 1])

        return ans