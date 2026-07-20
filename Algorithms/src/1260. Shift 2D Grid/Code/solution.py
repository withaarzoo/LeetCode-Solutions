class Solution:
    def shiftGrid(self, grid: List[List[int]], k: int) -> List[List[int]]:

        # Get the grid dimensions
        m = len(grid)
        n = len(grid[0])

        # Total number of elements
        total = m * n

        # Ignore unnecessary complete rotations
        k %= total

        # Create the answer grid
        ans = [[0] * n for _ in range(m)]

        # Move every element to its final position
        for i in range(m):
            for j in range(n):

                # Flatten the current position
                old_index = i * n + j

                # Compute the shifted position
                new_index = (old_index + k) % total

                # Convert back to row and column
                new_row = new_index // n
                new_col = new_index % n

                # Place the value
                ans[new_row][new_col] = grid[i][j]

        return ans