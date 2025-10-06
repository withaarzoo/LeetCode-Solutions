import heapq
from typing import List

class Solution:
    def swimInWater(self, grid: List[List[int]]) -> int:
        n = len(grid)
        visited = [[False] * n for _ in range(n)]
        # heap stores tuples (time, r, c)
        heap = [(grid[0][0], 0, 0)]
        heapq.heapify(heap)
        dirs = [(1,0),(-1,0),(0,1),(0,-1)]

        while heap:
            t, r, c = heapq.heappop(heap)
            if visited[r][c]:
                continue
            visited[r][c] = True
            # first time we pop target is the minimum possible time
            if r == n - 1 and c == n - 1:
                return t
            for dr, dc in dirs:
                nr, nc = r + dr, c + dc
                if 0 <= nr < n and 0 <= nc < n and not visited[nr][nc]:
                    nt = max(t, grid[nr][nc])
                    heapq.heappush(heap, (nt, nr, nc))
        return -1
