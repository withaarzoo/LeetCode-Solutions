from collections import deque
from typing import List

class Solution:
    def findSafeWalk(self, grid: List[List[int]], health: int) -> bool:
        m, n = len(grid), len(grid[0])

        # Store the minimum health lost to reach every cell
        dist = [[float("inf")] * n for _ in range(m)]

        # Deque used by 0-1 BFS
        dq = deque()

        # Starting cost includes the starting cell
        dist[0][0] = grid[0][0]
        dq.appendleft((0, 0))

        directions = [(-1, 0), (1, 0), (0, -1), (0, 1)]

        while dq:
            x, y = dq.popleft()

            # Visit all neighboring cells
            for dx, dy in directions:
                nx, ny = x + dx, y + dy

                # Ignore invalid positions
                if not (0 <= nx < m and 0 <= ny < n):
                    continue

                # Cost after entering the next cell
                new_cost = dist[x][y] + grid[nx][ny]

                # Update if this path is better
                if new_cost < dist[nx][ny]:
                    dist[nx][ny] = new_cost

                    # Weight 0 goes to the front
                    if grid[nx][ny] == 0:
                        dq.appendleft((nx, ny))
                    else:
                        dq.append((nx, ny))

        # Final health must remain positive
        return dist[m - 1][n - 1] < health