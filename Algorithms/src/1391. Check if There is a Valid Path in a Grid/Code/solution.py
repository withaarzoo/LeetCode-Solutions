from typing import List
from collections import deque

class Solution:
    def hasValidPath(self, grid: List[List[int]]) -> bool:
        m, n = len(grid), len(grid[0])

        # Directions: left, right, up, down
        dirs = [
            (0, -1),  # left
            (0, 1),   # right
            (-1, 0),  # up
            (1, 0)    # down
        ]

        # For each street type, which directions it supports.
        # 0 = left, 1 = right, 2 = up, 3 = down
        street_dirs = {
            1: [0, 1],  # left-right
            2: [2, 3],  # up-down
            3: [0, 3],  # left-down
            4: [1, 3],  # right-down
            5: [0, 2],  # left-up
            6: [1, 2],  # right-up
        }

        opposite = {0: 1, 1: 0, 2: 3, 3: 2}

        visited = [[False] * n for _ in range(m)]
        q = deque([(0, 0)])
        visited[0][0] = True

        while q:
            r, c = q.popleft()

            if r == m - 1 and c == n - 1:
                return True

            for d in street_dirs[grid[r][c]]:
                nr = r + dirs[d][0]
                nc = c + dirs[d][1]

                if nr < 0 or nr >= m or nc < 0 or nc >= n or visited[nr][nc]:
                    continue

                next_type = grid[nr][nc]

                # Check whether the next street connects back to the current cell
                if opposite[d] in street_dirs[next_type]:
                    visited[nr][nc] = True
                    q.append((nr, nc))

        return False