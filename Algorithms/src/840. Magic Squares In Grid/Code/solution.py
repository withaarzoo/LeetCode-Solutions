class Solution:
    def numMagicSquaresInside(self, grid: List[List[int]]) -> int:
        # Define a helper function to check if a 3x3 grid starting at (r, c) is a magic square
        def isMagicSquare(r: int, c: int) -> bool:
            # Create an array to track the occurrence of numbers 1 through 9
            vals = [0] * 10  # We use index 1-9, index 0 is unused

            # Iterate through the 3x3 grid
            for i in range(3):
                for j in range(3):
                    num = grid[r + i][c + j]  # Get the number in the current cell

                    # Check if the number is valid (between 1 and 9) and hasn't been used before
                    if num < 1 or num > 9 or vals[num]:
                        return False  # If invalid or duplicate, return False

                    vals[num] = 1  # Mark the number as used

            # Check if all rows, columns, and diagonals sum up to 15
            return (grid[r][c] + grid[r][c+1] + grid[r][c+2] == 15 and  # First row
                    grid[r+1][c] + grid[r+1][c+1] + grid[r+1][c+2] == 15 and  # Second row
                    grid[r+2][c] + grid[r+2][c+1] + grid[r+2][c+2] == 15 and  # Third row
                    grid[r][c] + grid[r+1][c] + grid[r+2][c] == 15 and  # First column
                    grid[r][c+1] + grid[r+1][c+1] + grid[r+2][c+1] == 15 and  # Second column
                    grid[r][c+2] + grid[r+1][c+2] + grid[r+2][c+2] == 15 and  # Third column
                    grid[r][c] + grid[r+1][c+1] + grid[r+2][c+2] == 15 and  # Main diagonal (top-left to bottom-right)
                    grid[r][c+2] + grid[r+1][c+1] + grid[r+2][c] == 15)  # Anti-diagonal (top-right to bottom-left)

        rows, cols = len(grid), len(grid[0])  # Get the number of rows and columns in the grid
        count = 0  # Initialize a counter for magic squares

        # Loop through the grid, checking each possible 3x3 sub-grid
        for i in range(rows - 2):
            for j in range(cols - 2):
                if isMagicSquare(i, j):  # Check if the 3x3 grid starting at (i, j) is a magic square
                    count += 1  # Increment the counter if a magic square is found

        return count  # Return the total number of magic squares found
