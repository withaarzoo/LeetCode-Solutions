import heapq
from typing import List

class Solution:
    def minCost(self, n: int, edges: List[List[int]]) -> int:
        graph = [[] for _ in range(n)]

        for u, v, w in edges:
            graph[u].append((v, w))
            graph[v].append((u, 2 * w))

        dist = [float('inf')] * n
        dist[0] = 0
        pq = [(0, 0)]

        while pq:
            cost, node = heapq.heappop(pq)

            if cost > dist[node]:
                continue

            for nxt, w in graph[node]:
                if dist[nxt] > cost + w:
                    dist[nxt] = cost + w
                    heapq.heappush(pq, (dist[nxt], nxt))

        return -1 if dist[n - 1] == float('inf') else dist[n - 1]
