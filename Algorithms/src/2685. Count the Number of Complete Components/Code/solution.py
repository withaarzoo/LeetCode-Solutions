class Solution:
    def countCompleteComponents(self, n: int, edges: List[List[int]]) -> int:

        # Build adjacency list
        graph = [[] for _ in range(n)]

        for u, v in edges:
            graph[u].append(v)
            graph[v].append(u)

        # Keep track of visited nodes
        visited = [False] * n

        # DFS to explore one connected component
        def dfs(node):

            # Mark current node
            visited[node] = True

            # Count current vertex
            vertices = 1

            # Add current node's degree
            degree_sum = len(graph[node])

            # Visit all neighbors
            for nxt in graph[node]:
                if not visited[nxt]:
                    v, d = dfs(nxt)
                    vertices += v
                    degree_sum += d

            return vertices, degree_sum

        answer = 0

        # Process every component
        for i in range(n):

            if visited[i]:
                continue

            vertices, degree_sum = dfs(i)

            # Every edge appears twice
            edge_count = degree_sum // 2

            # Check if this component is complete
            if edge_count == vertices * (vertices - 1) // 2:
                answer += 1

        return answer