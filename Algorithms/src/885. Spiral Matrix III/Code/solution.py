from typing import List

class Solution:
    def spiralMatrixIII(self, rows: int, cols: int, rStart: int, cStart: int) -> List[List[int]]:
        # Initialize the result list to store the coordinates of cells in spiral order
        result = []
        
        # Define the four possible movement directions: right, down, left, and up
        directions = [(0, 1), (1, 0), (0, -1), (-1, 0)]
        
        # Start with one step to the right
        steps = 1
        
        # Initialize the direction index to point to the first direction (right)
        d = 0
        
        # Set the starting position
        r, c = rStart, cStart
        
        # Add the starting position to the result
        result.append([r, c])

        # Continue the loop until all cells are visited
        while len(result) < rows * cols:
            # Each direction is taken twice for each 'step' increment
            for _ in range(2):
                # Move 'steps' number of times in the current direction
                for _ in range(steps):
                    # Update the row and column indices based on the current direction
                    r += directions[d][0]
                    c += directions[d][1]
                    # Check if the new position is within the bounds of the matrix
                    if 0 <= r < rows and 0 <= c < cols:
                        # If it is, add it to the result list
                        result.append([r, c])
                # Change to the next direction (right -> down -> left -> up -> right -> ...)
                d = (d + 1) % 4
            # After moving in two directions, increase the step count
            steps += 1

        # Return the result list containing the coordinates of the cells in spiral order
        return result
