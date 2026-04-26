from typing import List

class Solution:
    def containsCycle(self, grid: List[List[str]]) -> bool:
        m, n = len(grid), len(grid[0])
        visited = [[False] * n for _ in range(m)]

        dirs = [(1, 0), (-1, 0), (0, 1), (0, -1)]

        for r in range(m):
            for c in range(n):
                if visited[r][c]:
                    continue

                # stack item: (row, col, parent_row, parent_col)
                stack = [(r, c, -1, -1)]
                visited[r][c] = True

                while stack:
                    cr, cc, pr, pc = stack.pop()

                    for dr, dc in dirs:
                        nr, nc = cr + dr, cc + dc

                        if nr < 0 or nr >= m or nc < 0 or nc >= n:
                            continue
                        if grid[nr][nc] != grid[cr][cc]:
                            continue
                        if nr == pr and nc == pc:
                            continue

                        if visited[nr][nc]:
                            return True

                        visited[nr][nc] = True
                        stack.append((nr, nc, cr, cc))

        return False