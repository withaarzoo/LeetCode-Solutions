class Solution:
    def reverseSubmatrix(self, grid, x, y, k):
        # Loop half rows
        for i in range(k // 2):
            top = x + i
            bottom = x + k - 1 - i

            # Swap columns
            for j in range(k):
                grid[top][y + j], grid[bottom][y + j] = grid[bottom][y + j], grid[top][y + j]

        return grid