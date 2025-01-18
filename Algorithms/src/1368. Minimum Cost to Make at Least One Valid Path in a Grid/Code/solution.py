from collections import deque

class Solution:
    def minCost(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        directions = [(0, 1), (0, -1), (1, 0), (-1, 0)]
        cost = [[float('inf')] * n for _ in range(m)]
        dq = deque([(0, 0)])
        cost[0][0] = 0
        
        while dq:
            x, y = dq.popleft()
            for i, (dx, dy) in enumerate(directions):
                nx, ny = x + dx, y + dy
                new_cost = cost[x][y] + (grid[x][y] != i + 1)
                
                if 0 <= nx < m and 0 <= ny < n and new_cost < cost[nx][ny]:
                    cost[nx][ny] = new_cost
                    if grid[x][y] == i + 1:
                        dq.appendleft((nx, ny))
                    else:
                        dq.append((nx, ny))
        return cost[m - 1][n - 1]
