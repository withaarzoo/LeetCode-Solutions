class Solution:
    def findMaxFish(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        visited = [[False] * n for _ in range(m)]
        max_fish = 0

        def dfs(r, c):
            if r < 0 or c < 0 or r >= m or c >= n or visited[r][c] or grid[r][c] == 0:
                return 0
            visited[r][c] = True
            fish = grid[r][c]
            fish += dfs(r + 1, c)
            fish += dfs(r - 1, c)
            fish += dfs(r, c + 1)
            fish += dfs(r, c - 1)
            return fish

        for i in range(m):
            for j in range(n):
                if not visited[i][j] and grid[i][j] > 0:
                    max_fish = max(max_fish, dfs(i, j))
        return max_fish
