class Solution:
    def numMagicSquaresInside(self, grid):
        rows, cols = len(grid), len(grid[0])
        count = 0

        for i in range(rows - 2):
            for j in range(cols - 2):
                if self.isMagic(grid, i, j):
                    count += 1
        return count

    def isMagic(self, g, r, c):
        if g[r + 1][c + 1] != 5:
            return False

        seen = set()

        for i in range(r, r + 3):
            for j in range(c, c + 3):
                val = g[i][j]
                if val < 1 or val > 9 or val in seen:
                    return False
                seen.add(val)

        for i in range(3):
            if sum(g[r + i][c:c + 3]) != 15:
                return False
            if g[r][c + i] + g[r + 1][c + i] + g[r + 2][c + i] != 15:
                return False

        if g[r][c] + g[r + 1][c + 1] + g[r + 2][c + 2] != 15:
            return False
        if g[r][c + 2] + g[r + 1][c + 1] + g[r + 2][c] != 15:
            return False

        return True
