from collections import deque
from typing import List

class Solution:
    def minDays(self, grid: List[List[int]]) -> int:
        # Step 1: Check if the grid is already disconnected.
        if self.isDisconnected(grid):
            return 0  # If the grid is disconnected, no need to remove any cell.

        m, n = len(grid), len(grid[0])  # Get the dimensions of the grid.

        # Step 2: Try removing one cell at a time.
        for i in range(m):
            for j in range(n):
                if grid[i][j] == 1:  # If the cell contains land (1),
                    grid[i][j] = 0  # Temporarily remove this cell.
                    if self.isDisconnected(grid):
                        return 1  # If the grid becomes disconnected, return 1.
                    grid[i][j] = 1  # Restore the cell if it didn't disconnect the grid.

        # Step 3: Try removing two cells.
        for i in range(m):
            for j in range(n):
                if grid[i][j] == 1:  # If the cell contains land (1),
                    grid[i][j] = 0  # Temporarily remove this cell.
                    for x in range(m):
                        for y in range(n):
                            if grid[x][y] == 1:  # Find another land cell.
                                grid[x][y] = 0  # Temporarily remove the second cell.
                                if self.isDisconnected(grid):
                                    return 2  # If the grid becomes disconnected, return 2.
                                grid[x][y] = 1  # Restore the second cell if it didn't disconnect the grid.
                    grid[i][j] = 1  # Restore the first cell.

        return 2  # If no solution found, return 2 as the grid will be disconnected by removing two cells.

    def isDisconnected(self, grid: List[List[int]]) -> bool:
        m, n = len(grid), len(grid[0])  # Get the dimensions of the grid.
        visited = [[0] * n for _ in range(m)]  # Initialize a visited matrix to track visited cells.

        land_count = 0  # Count the number of separate land areas.
        for i in range(m):
            for j in range(n):
                if grid[i][j] == 1:  # If the cell contains land (1),
                    land_count += 1  # Increment the land area count.
                    if not visited[i][j]:  # If this land cell is not visited,
                        if land_count > 1:  # If more than one land area is found,
                            return True  # The grid is already disconnected.
                        self.bfs(grid, visited, i, j)  # Perform BFS to mark all connected land cells as visited.

        return land_count == 0  # Return true if the grid has no land or a single connected land area.

    def bfs(self, grid: List[List[int]], visited: List[List[int]], i: int, j: int) -> None:
        m, n = len(grid), len(grid[0])  # Get the dimensions of the grid.
        queue = deque([(i, j)])  # Initialize a queue for BFS with the starting cell.
        visited[i][j] = 1  # Mark the starting cell as visited.

        # Define the directions for moving up, down, left, and right.
        dirX = [-1, 1, 0, 0]
        dirY = [0, 0, -1, 1]

        # Perform BFS to visit all connected land cells.
        while queue:
            x, y = queue.popleft()  # Get the current cell from the queue.

            # Explore the 4 possible directions.
            for d in range(4):
                newX = x + dirX[d]
                newY = y + dirY[d]
                # Check if the new position is within bounds and is a land cell that hasn't been visited.
                if 0 <= newX < m and 0 <= newY < n and grid[newX][newY] == 1 and not visited[newX][newY]:
                    visited[newX][newY] = 1  # Mark the new cell as visited.
                    queue.append((newX, newY))  # Add the new cell to the queue for further exploration.
