from collections import deque
from typing import List

class Solution:
    def pacificAtlantic(self, heights: List[List[int]]) -> List[List[int]]:
        if not heights or not heights[0]:
            return []
        m, n = len(heights), len(heights[0])

        pac = [[False]*n for _ in range(m)]
        atl = [[False]*n for _ in range(m)]

        def bfs(starts, visited):
            q = deque(starts)
            while q:
                r, c = q.popleft()
                for dr, dc in ((1,0),(-1,0),(0,1),(0,-1)):
                    nr, nc = r + dr, c + dc
                    if nr < 0 or nr >= m or nc < 0 or nc >= n:
                        continue
                    if visited[nr][nc]:
                        continue
                    if heights[nr][nc] < heights[r][c]:
                        continue
                    visited[nr][nc] = True
                    q.append((nr, nc))

        # Pacific: top row and left column
        pac_starts = [(0, j) for j in range(n)]
        pac_starts += [(i, 0) for i in range(1, m)]
        for r, c in pac_starts:
            pac[r][c] = True
        bfs(pac_starts, pac)

        # Atlantic: bottom row and right column
        atl_starts = [(m - 1, j) for j in range(n)]
        atl_starts += [(i, n - 1) for i in range(0, m - 1)]
        for r, c in atl_starts:
            atl[r][c] = True
        bfs(atl_starts, atl)

        ans = []
        for i in range(m):
            for j in range(n):
                if pac[i][j] and atl[i][j]:
                    ans.append([i, j])
        return ans
