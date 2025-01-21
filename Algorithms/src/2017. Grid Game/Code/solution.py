class Solution:
    def gridGame(self, grid: List[List[int]]) -> int:
        n = len(grid[0])
        
        # Calculate suffix sum for the top row
        top_suffix = [0] * n
        top_suffix[n - 1] = grid[0][n - 1]
        for i in range(n - 2, -1, -1):
            top_suffix[i] = top_suffix[i + 1] + grid[0][i]
        
        # Calculate prefix sum for the bottom row
        bottom_prefix = [0] * n
        bottom_prefix[0] = grid[1][0]
        for i in range(1, n):
            bottom_prefix[i] = bottom_prefix[i - 1] + grid[1][i]
        
        # Find the minimum maximum points Robot 2 can collect
        result = float('inf')
        for i in range(n):
            top = top_suffix[i + 1] if i + 1 < n else 0
            bottom = bottom_prefix[i - 1] if i > 0 else 0
            result = min(result, max(top, bottom))
        
        return result
