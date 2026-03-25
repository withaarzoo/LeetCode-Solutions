class Solution:
    def canPartitionGrid(self, grid):
        m, n = len(grid), len(grid[0])
        
        # Step 1: Total sum
        total = sum(sum(row) for row in grid)
        
        # Step 2: Odd check
        if total % 2 != 0:
            return False
        
        target = total // 2
        
        # Step 3: Horizontal cut
        row_sum = 0
        for i in range(m - 1):
            row_sum += sum(grid[i])
            if row_sum == target:
                return True
        
        # Step 4: Column sums
        col_sum = [0] * n
        for j in range(n):
            for i in range(m):
                col_sum[j] += grid[i][j]
        
        # Step 5: Vertical cut
        curr = 0
        for j in range(n - 1):
            curr += col_sum[j]
            if curr == target:
                return True
        
        return False