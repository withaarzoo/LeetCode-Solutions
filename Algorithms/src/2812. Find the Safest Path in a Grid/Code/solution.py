from collections import deque

class Solution:
    def maximumSafenessFactor(self, grid: List[List[int]]) -> int:

        n = len(grid)

        # Distance from the nearest thief
        dist = [[-1] * n for _ in range(n)]

        q = deque()

        # Push every thief into the queue
        for i in range(n):
            for j in range(n):
                if grid[i][j] == 1:
                    dist[i][j] = 0
                    q.append((i, j))

        directions = [(-1,0),(1,0),(0,-1),(0,1)]

        # Multi-source BFS
        while q:
            x, y = q.popleft()

            for dx, dy in directions:
                nx, ny = x + dx, y + dy

                if 0 <= nx < n and 0 <= ny < n and dist[nx][ny] == -1:
                    dist[nx][ny] = dist[x][y] + 1
                    q.append((nx, ny))

        # Check whether a path exists
        def canReach(limit):

            if dist[0][0] < limit or dist[n-1][n-1] < limit:
                return False

            vis = [[False] * n for _ in range(n)]
            bfs = deque([(0, 0)])
            vis[0][0] = True

            while bfs:

                x, y = bfs.popleft()

                if x == n - 1 and y == n - 1:
                    return True

                for dx, dy in directions:

                    nx, ny = x + dx, y + dy

                    if 0 <= nx < n and 0 <= ny < n:
                        if not vis[nx][ny] and dist[nx][ny] >= limit:
                            vis[nx][ny] = True
                            bfs.append((nx, ny))

            return False

        left = 0
        right = 2 * n
        ans = 0

        # Binary search on the answer
        while left <= right:

            mid = (left + right) // 2

            if canReach(mid):
                ans = mid
                left = mid + 1
            else:
                right = mid - 1

        return ans