class Solution:
    def shortestPath(self, n: int, edges: List[List[int]], labels: str, k: int) -> int:
        g = [[] for _ in range(n)]
        for u, v, w in edges:
            g[u].append((v, w))
        INF = 10**9
        dist = [[INF] * (k + 1) for _ in range(n)]
        pq = [(0, 0, 1)]
        dist[0][1] = 0
        while pq:
            d, u, cnt = heapq.heappop(pq)
            if d > dist[u][cnt]:
                continue
            if u == n - 1:
                return d
            for v, w in g[u]:
                ncnt = cnt + 1 if labels[u] == labels[v] else 1
                if ncnt > k:
                    continue
                if d + w < dist[v][ncnt]:
                    dist[v][ncnt] = d + w
                    heapq.heappush(pq, (dist[v][ncnt], v, ncnt))
        return -1