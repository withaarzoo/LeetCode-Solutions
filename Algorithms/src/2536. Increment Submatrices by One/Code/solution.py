class Solution:
    def rangeAddQueries(self, n: int, queries: List[List[int]]) -> List[List[int]]:
        # diff is (n+1) x (n+1)
        diff = [[0] * (n+1) for _ in range(n+1)]
        
        # Apply corner updates for each query
        for r1, c1, r2, c2 in queries:
            diff[r1][c1] += 1
            diff[r1][c2 + 1] -= 1
            diff[r2 + 1][c1] -= 1
            diff[r2 + 1][c2 + 1] += 1
        
        # Convert diff to final matrix via 2D prefix sum
        res = [[0] * n for _ in range(n)]
        for i in range(n):
            for j in range(n):
                up = diff[i-1][j] if i > 0 else 0
                left = diff[i][j-1] if j > 0 else 0
                diag = diff[i-1][j-1] if i > 0 and j > 0 else 0
                diff[i][j] = diff[i][j] + up + left - diag
                res[i][j] = diff[i][j]
        return res
