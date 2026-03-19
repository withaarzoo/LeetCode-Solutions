class Solution:
    def numberOfSubmatrices(self, grid):
        n, m = len(grid), len(grid[0])
        
        sum_arr = [[0] * (m + 1) for _ in range(2)]
        countX = [[0] * (m + 1) for _ in range(2)]
        
        ans = 0
        
        for i in range(n):
            cur = i % 2
            prev = 1 - cur
            
            for j in range(m):
                val = 1 if grid[i][j] == 'X' else (-1 if grid[i][j] == 'Y' else 0)
                isX = 1 if grid[i][j] == 'X' else 0
                
                sum_arr[cur][j + 1] = val \
                    + sum_arr[cur][j] \
                    + sum_arr[prev][j + 1] \
                    - sum_arr[prev][j]
                
                countX[cur][j + 1] = isX \
                    + countX[cur][j] \
                    + countX[prev][j + 1] \
                    - countX[prev][j]
                
                if sum_arr[cur][j + 1] == 0 and countX[cur][j + 1] > 0:
                    ans += 1
        
        return ans