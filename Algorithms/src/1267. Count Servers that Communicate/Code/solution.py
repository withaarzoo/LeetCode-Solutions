class Solution:
    def countServers(self, grid: List[List[int]]) -> int:
        rows, cols = len(grid), len(grid[0])
        rowCount = [0] * rows
        colCount = [0] * cols
        
        # First pass: Count servers in each row and column
        for i in range(rows):
            for j in range(cols):
                if grid[i][j] == 1:
                    rowCount[i] += 1
                    colCount[j] += 1
        
        # Second pass: Count communicable servers
        count = 0
        for i in range(rows):
            for j in range(cols):
                if grid[i][j] == 1 and (rowCount[i] > 1 or colCount[j] > 1):
                    count += 1
        
        return count
