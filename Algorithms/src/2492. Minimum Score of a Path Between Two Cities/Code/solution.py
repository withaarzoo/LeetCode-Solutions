class Solution:
    def minScore(self, n: int, roads: List[List[int]]) -> int:
        # graph[city] stores pairs of (neighbor, road distance).
        graph = [[] for _ in range(n + 1)]

        # Build the undirected graph by storing each road in both directions.
        for a, b, distance in roads:
            graph[a].append((b, distance))
            graph[b].append((a, distance))

        # visited prevents the same city from being processed repeatedly.
        visited = [False] * (n + 1)

        # I use a deque for efficient BFS operations.
        queue = deque([1])
        visited[1] = True

        # Start with an infinitely large value.
        answer = float("inf")

        # Visit every city connected to city 1.
        while queue:
            city = queue.popleft()

            # Check every road connected to the current city.
            for next_city, distance in graph[city]:
                # Keep the smallest road distance found in this component.
                answer = min(answer, distance)

                # Visit the neighboring city only if it is still unvisited.
                if not visited[next_city]:
                    visited[next_city] = True
                    queue.append(next_city)

        # Return the smallest road distance in city 1's component.
        return answer