class Solution:
    def firstCompleteIndex(self, arr: List[int], mat: List[List[int]]) -> int:
        m, n = len(mat), len(mat[0])
        position = {}
        rowCount = [0] * m
        colCount = [0] * n

        # Map matrix values to their positions
        for i in range(m):
            for j in range(n):
                position[mat[i][j]] = (i, j)

        # Iterate through the array and simulate painting
        for i, val in enumerate(arr):
            row, col = position[val]
            rowCount[row] += 1
            colCount[col] += 1

            if rowCount[row] == n or colCount[col] == m:
                return i
        return -1
