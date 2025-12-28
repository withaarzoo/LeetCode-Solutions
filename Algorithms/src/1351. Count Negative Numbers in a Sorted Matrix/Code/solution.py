class Solution:
    def countNegatives(self, grid: List[List[int]]) -> int:
        rows = len(grid)
        cols = len(grid[0])

        r = 0
        c = cols - 1
        count = 0

        # Start from top-right corner
        while r < rows and c >= 0:
            if grid[r][c] < 0:
                count += (rows - r)
                c -= 1  # move left
            else:
                r += 1  # move down

        return count
