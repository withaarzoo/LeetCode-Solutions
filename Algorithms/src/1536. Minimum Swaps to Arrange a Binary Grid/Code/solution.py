class Solution:
    def minSwaps(self, grid: List[List[int]]) -> int:
        n = len(grid)
        
        # Count trailing zeros
        trailing = []
        for i in range(n):
            count = 0
            for j in range(n - 1, -1, -1):
                if grid[i][j] == 0:
                    count += 1
                else:
                    break
            trailing.append(count)
        
        swaps = 0
        
        for i in range(n):
            required = n - 1 - i
            j = i
            
            while j < n and trailing[j] < required:
                j += 1
            
            if j == n:
                return -1
            
            while j > i:
                trailing[j], trailing[j - 1] = trailing[j - 1], trailing[j]
                swaps += 1
                j -= 1
        
        return swaps